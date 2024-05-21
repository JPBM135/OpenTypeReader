package font

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func ParseFont(reader *bytes.Buffer) error {
	result, err := ParseTableDirectory(reader)

	if err != nil {
		return err
	}

	fmt.Printf("\n\n")
	resultJson, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	os.WriteFile("font.json", resultJson, 0600)
	return nil
}
