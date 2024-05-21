package tableparser

import (
	"bytes"

	"jpbm135.open-type-reader/pkg/readers"
	"jpbm135.open-type-reader/pkg/structs"
	"jpbm135.open-type-reader/pkg/utils"
)

// https://learn.microsoft.com/en-us/typography/opentype/spec/cmap

const (
	PLATFORM_ID_UNICODE   = 0
	PLATFORM_ID_MACINTOSH = 1
	PLATFORM_ID_ISO       = 2
	PLATFORM_ID_WINDOWS   = 3
	PLATFORM_ID_CUSTOM    = 4
)

func ParseCmapTable(buffer *bytes.Buffer, cmapTable structs.FontTableRecord) (structs.CmapTable, error) {
	cmapReader := utils.ReaderFromBuffer(buffer)
	cmapReader.Seek(int64(cmapTable.Offset), 0)

	version := readers.ReadUInt16(cmapReader)
	numTables := readers.ReadUInt16(cmapReader)

	var encodingRecords []structs.CmapEncodingRecord = make([]structs.CmapEncodingRecord, 0)
	for i := 0; i < int(numTables); i++ {
		encodingRecord := structs.CmapEncodingRecord{
			PlatformID:     readers.ReadUInt16(cmapReader),
			EncodingID:     readers.ReadUInt16(cmapReader),
			SubtableOffset: readers.ReadUInt32(cmapReader),
		}

		encodingRecords = append(encodingRecords, encodingRecord)
	}

	for i := 0; i < len(encodingRecords); i++ {
		encodingRecord := encodingRecords[i]
		parsedSubtable := ParseCmapSubtable(buffer, int64(cmapTable.Offset+encodingRecord.SubtableOffset))

		encodingRecords[i].Subtable = parsedSubtable
	}

	cmap := structs.CmapTable{
		Version:         version,
		NumTables:       numTables,
		EncodingRecords: encodingRecords,
	}

	return cmap, nil
}

func ParseCmapSubtable(buffer *bytes.Buffer, offset int64) structs.CmapSubtable {
	cmapReader := utils.ReaderFromBuffer(buffer)
	cmapReader.Seek(offset, 0)

	format := readers.ReadUInt16(cmapReader)

	switch format {
	case 0:
		return parseCmapFormat0Subtable(cmapReader)
	case 4:
		return parseCmapFormat4Subtable(cmapReader)
	default:
		return structs.CmapEncodingRecordSubtableFormatUnknown{
			Format:   format,
			Length:   readers.ReadUInt16(cmapReader),
			Language: readers.ReadUInt16(cmapReader),
		}
	}
}

func parseCmapFormat0Subtable(reader *bytes.Reader) structs.CmapEncodingRecordSubtableFormat0 {
	length := readers.ReadUInt16(reader)
	language := readers.ReadUInt16(reader)

	var glyphIdArray [256]uint8
	for i := 0; i < 256; i++ {
		glyphIdArray[i] = readers.ReadUInt8(reader)
	}

	recordSubtable := structs.CmapEncodingRecordSubtableFormat0{
		Format:       0,
		Length:       length,
		Language:     language,
		GlyphIdArray: glyphIdArray,
	}

	return recordSubtable
}

func parseCmapFormat4Subtable(reader *bytes.Reader) structs.CmapEncodingRecordSubtableFormat4 {
	length := readers.ReadUInt16(reader)
	language := readers.ReadUInt16(reader)
	segCountX2 := readers.ReadUInt16(reader)
	segCount := segCountX2 / 2
	searchRange := readers.ReadUInt16(reader)
	entrySelector := readers.ReadUInt16(reader)
	rangeShift := readers.ReadUInt16(reader)

	var endCode []uint16 = make([]uint16, segCount)
	for i := 0; i < int(segCount); i++ {
		endCode = append(endCode, readers.ReadUInt16(reader))
	}

	reservedPad := readers.ReadUInt16(reader)

	var startCode []uint16 = make([]uint16, segCount)
	for i := 0; i < int(segCount); i++ {
		startCode = append(startCode, readers.ReadUInt16(reader))
	}

	var idDelta []int16 = make([]int16, segCount)
	for i := 0; i < int(segCount); i++ {
		idDelta = append(idDelta, readers.ReadInt16(reader))
	}

	var idRangeOffsets []uint16 = make([]uint16, segCount)
	for i := 0; i < int(segCount); i++ {
		idRangeOffsets = append(idRangeOffsets, readers.ReadUInt16(reader))
	}

	var glyphIdArray []uint16 = make([]uint16, segCount)
	for i := 0; i < int(segCount); i++ {
		glyphIdArray = append(glyphIdArray, readers.ReadUInt16(reader))
	}

	recordSubtable := structs.CmapEncodingRecordSubtableFormat4{
		Format:         0,
		Length:         length,
		Language:       language,
		SegCountX2:     segCountX2,
		SearchRange:    searchRange,
		EntrySelector:  entrySelector,
		RangeShift:     rangeShift,
		EndCode:        endCode,
		ReservedPad:    reservedPad,
		StartCode:      startCode,
		IdDelta:        idDelta,
		IdRangeOffsets: idRangeOffsets,
		GlyphIdArray:   glyphIdArray,
	}

	return recordSubtable
}
