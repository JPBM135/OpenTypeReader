package structs

// https://learn.microsoft.com/en-us/typography/opentype/spec/otff#table-directory

// FontTableDirectory represents the table directory in the font file
type FontTableDirectory struct {
	SfntVersion   uint32                     `json:"sfntVersion"`
	NumTables     uint16                     `json:"numTables"`
	SearchRange   uint16                     `json:"searchRange"`
	EntrySelector uint16                     `json:"entrySelector"`
	RangeShift    uint16                     `json:"rangeShift"`
	Tables        map[string]FontTableRecord `json:"tables"`
}

// FontTableRecord represents a record in the table directory
type FontTableRecord struct {
	TableTag string `json:"tableTag"`
	CheckSum uint32 `json:"checkSum"`
	Offset   uint32 `json:"offset"`
	Length   uint32 `json:"length"`
	Parsed   any    `json:"parsed"`
}
