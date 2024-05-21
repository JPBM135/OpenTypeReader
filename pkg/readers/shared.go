package readers

import "io"

func ReadByte(reader io.Reader, byteAmount int) []byte {
	bytes := make([]byte, byteAmount)

	_, err := reader.Read(bytes)
	if err != nil {
		panic(err)
	}

	if len(bytes) < byteAmount {
		panic(io.ErrUnexpectedEOF)
	}

	return bytes
}
