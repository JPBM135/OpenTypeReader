package readers

import (
	"io"
)

// https://learn.microsoft.com/en-us/typography/opentype/spec/otff#data-types

func ReadTag(reader io.Reader) string {
	bytes := make([]byte, 4)

	_, err := reader.Read(bytes)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
