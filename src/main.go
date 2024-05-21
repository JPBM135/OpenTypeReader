package main

import (
	"bytes"
	"fmt"
	"os"

	"jpbm135.open-type-reader/pkg/font"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fontFile := dir + "/data/Roboto-Regular.ttf"
	fontFile := dir + "/data/Arial.ttf"
	fmt.Println("Font file path:", fontFile)

	fontBytes, err := os.ReadFile(fontFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader := bytes.NewBuffer(fontBytes)

	err = font.ParseFont(reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Font file read successfully", len(fontBytes), "bytes")
}
