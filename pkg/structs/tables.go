package structs

// CmapTable represents the 'cmap' table
type CmapTable struct {
	Version         uint16               `json:"version"`
	NumTables       uint16               `json:"numTables"`
	EncodingRecords []CmapEncodingRecord `json:"encodingRecords"`
}

// CmapEncodingRecord represents the encoding record in the 'cmap' table
type CmapEncodingRecord struct {
	PlatformID uint16 `json:"platformID"`
	EncodingID uint16 `json:"encodingID"`
	// Offset from beginning of the 'cmap' table
	SubtableOffset uint32       `json:"subtableOffset"`
	Subtable       CmapSubtable `json:"subtable"`
}

type CmapSubtable interface {
	// This is a marker interface, so it doesn't need any methods
}

type CmapEncodingRecordSubtableFormatUnknown struct {
	Format   uint16 `json:"format"`
	Length   uint16 `json:"length"`
	Language uint16 `json:"language"`
}

// https://learn.microsoft.com/en-us/typography/opentype/spec/cmap#format-0-byte-encoding-table
type CmapEncodingRecordSubtableFormat0 struct {
	Format       uint16     `json:"format"`
	Length       uint16     `json:"length"`
	Language     uint16     `json:"language"`
	GlyphIdArray [256]uint8 `json:"glyphIdArray"`
}

// https://learn.microsoft.com/en-us/typography/opentype/spec/cmap#format-4-segment-mapping-to-delta-values
type CmapEncodingRecordSubtableFormat4 struct {
	Format         uint16   `json:"format"`
	Length         uint16   `json:"length"`
	Language       uint16   `json:"language"`
	SegCountX2     uint16   `json:"segCountX2"`
	SearchRange    uint16   `json:"searchRange"`
	EntrySelector  uint16   `json:"entrySelector"`
	RangeShift     uint16   `json:"rangeShift"`
	EndCode        []uint16 `json:"endCode"`
	ReservedPad    uint16   `json:"reservedPad"`
	StartCode      []uint16 `json:"startCode"`
	IdDelta        []int16  `json:"idDelta"`
	IdRangeOffsets []uint16 `json:"idRangeOffsets"`
	GlyphIdArray   []uint16 `json:"glyphIdArray"`
}

// https://learn.microsoft.com/en-us/typography/opentype/spec/cmap#format-10-trimmed-array

type CmapEncodingRecordSubtableFormat10 struct {
	Format        uint16   `json:"format"`
	Length        uint16   `json:"length"`
	Language      uint16   `json:"language"`
	StartCharCode uint32   `json:"startCharCode"`
	NumChars      uint32   `json:"numChars"`
	GlyphIdArray  []uint16 `json:"glyphIdArray"`
}

// https://learn.microsoft.com/en-us/typography/opentype/spec/cmap#format-12-segmented-coverage

type CmapEncodingRecordSubtableFormat12 struct {
	Format    uint16                                    `json:"format"`
	Length    uint16                                    `json:"length"`
	Language  uint16                                    `json:"language"`
	NumGroups uint32                                    `json:"numGroups"`
	Groups    []CmapEncodingRecordSubtableFormat12Group `json:"groups"`
}

type CmapEncodingRecordSubtableFormat12Group struct {
	StartCharCode uint32 `json:"startCharCode"`
	EndCharCode   uint32 `json:"endCharCode"`
	StartGlyphID  uint32 `json:"startGlyphID"`
}

// https://learn.microsoft.com/en-us/typography/opentype/spec/cmap#format-13-many-to-one-range-mappings
type CmapEncodingRecordSubtableFormat13 struct {
	Format    uint16                                    `json:"format"`
	Length    uint16                                    `json:"length"`
	Language  uint16                                    `json:"language"`
	NumGroups uint32                                    `json:"numGroups"`
	Groups    []CmapEncodingRecordSubtableFormat13Group `json:"groups"`
}

type CmapEncodingRecordSubtableFormat13Group struct {
	StartCharCode uint32 `json:"startCharCode"`
	EndCharCode   uint32 `json:"endCharCode"`
	GlyphID       uint32 `json:"glyphID"`
}

// https://learn.microsoft.com/en-us/typography/opentype/spec/cmap#format-14-unicode-variation-sequences

type CmapEncodingRecordSubtableFormat14 struct {
	Format                uint16                                                `json:"format"`
	Length                uint16                                                `json:"length"`
	NumVarSelectorRecords uint32                                                `json:"numVarSelectorRecords"`
	VarSelectorRecords    []CmapEncodingRecordSubtableFormat14VarSelectorRecord `json:"varSelectorRecords"`
}

type CmapEncodingRecordSubtableFormat14VarSelectorRecord struct {
	// VarSelector is a 24-bit value (how can we represent this in Go? Will use uint32 for now)
	VarSelector         uint32 `json:"varSelector"`
	DefaultUVSOffset    uint32 `json:"defaultUVSOffset"`
	NonDefaultUVSOffset uint32 `json:"nonDefaultUVSOffset"`
}
