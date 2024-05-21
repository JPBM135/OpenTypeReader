package readers

import (
	"encoding/binary"
	"io"
)

// https://learn.microsoft.com/en-us/typography/opentype/spec/otff#data-types

func ReadUInt32(reader io.Reader) uint32 {
	bytes := ReadByte(reader, 4)
	val := binary.BigEndian.Uint32(bytes)

	return val
}

func ReadUInt24(reader io.Reader) uint32 {
	bytes := ReadByte(reader, 3)

	val := (uint32(bytes[0]) | (uint32(bytes[1]) << 8) | (uint32(bytes[2]) << 16))

	return val
}

func ReadUInt16(reader io.Reader) uint16 {
	bytes := ReadByte(reader, 2)
	val := binary.BigEndian.Uint16(bytes)

	return val
}

func ReadUInt8(reader io.Reader) uint8 {
	bytes := ReadByte(reader, 1)
	return uint8(bytes[0])
}
