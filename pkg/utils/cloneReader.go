package utils

import (
	"bytes"
)

func ReaderFromBuffer(original *bytes.Buffer) *bytes.Reader {
	// Get the underlying byte slice
	underlyingBytes := original.Bytes()

	// Create a new bytes.Reader with the same byte slice
	clone := bytes.NewReader(underlyingBytes)

	return clone
}
