package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var InputFilename = "input.txt"

func InputReader() io.Reader {
	b, err := os.ReadFile(InputFilename)
	if err != nil {
		fmt.Printf("err reading file: %v\n", err)
		os.Exit(1)
	}

	return bytes.NewReader(b)
}
