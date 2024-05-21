package font

import (
	"bytes"
	"fmt"

	"jpbm135.open-type-reader/pkg/readers"
	"jpbm135.open-type-reader/pkg/structs"
	"jpbm135.open-type-reader/pkg/tableparser"
	"jpbm135.open-type-reader/pkg/utils"
)

func ParseTableDirectory(buffer *bytes.Buffer) (structs.FontTableDirectory, error) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered in ParseTableDirectory", r)
	// 	}
	// }()

	reader := utils.ReaderFromBuffer(buffer)

	// https://learn.microsoft.com/en-us/typography/opentype/spec/otff#table-directory
	sfntVersion := readers.ReadUInt32(reader)
	numberOfTables := readers.ReadUInt16(reader)
	searchRange := readers.ReadUInt16(reader)
	entrySelector := readers.ReadUInt16(reader)
	rangeShift := readers.ReadUInt16(reader)

	n := 1
	var tables map[string]structs.FontTableRecord = make(map[string]structs.FontTableRecord)
	for n <= int(numberOfTables) {
		table := readers.ReadTag(reader)
		fmt.Println(table)
		checksum := readers.ReadUInt32(reader)
		offset := readers.ReadUInt32(reader)
		length := readers.ReadUInt32(reader)

		tableRecord := structs.FontTableRecord{
			TableTag: table,
			CheckSum: checksum,
			Offset:   offset,
			Length:   length,
			Parsed:   nil,
		}

		parsed, err := tableparser.ParseTable(buffer, tableRecord)
		if err != nil {
			return structs.FontTableDirectory{}, nil
		}
		tableRecord.Parsed = parsed

		tables[table] = tableRecord

		n++
	}

	directory := structs.FontTableDirectory{
		SfntVersion:   sfntVersion,
		NumTables:     numberOfTables,
		SearchRange:   searchRange,
		EntrySelector: entrySelector,
		RangeShift:    rangeShift,
		Tables:        tables,
	}

	return directory, nil
}
