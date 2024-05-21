package readers

import (
	"io"
)

func ReadInt16(reader io.Reader) int16 {
	val := int16(ReadUInt16(reader))

	return val
}
