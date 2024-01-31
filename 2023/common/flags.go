package common

import (
	"encoding"
	"errors"
	"flag"
	"fmt"
	"os"
)

var inputFilename string

func init() {
	flag.StringVar(&inputFilename, "input", "input", "input filename")
}

func FileFlag(p encoding.TextUnmarshaler) error {
	if inputFilename == "" {
		return errors.New("blank filename")
	}

	b, err := os.ReadFile(inputFilename)
	if err != nil {
		return fmt.Errorf("unable to read: %w", err)
	}

	return p.UnmarshalText(b)
}
