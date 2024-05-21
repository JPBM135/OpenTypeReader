package tableparser

import (
	"bytes"

	"jpbm135.open-type-reader/pkg/structs"
)

func ParseTable(buffer *bytes.Buffer, record structs.FontTableRecord) (interface{}, error) {
	switch record.TableTag {
	case "cmap":
		return ParseCmapTable(buffer, record)
	default:
		return nil, nil
	}
}
