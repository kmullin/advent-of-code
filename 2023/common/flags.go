package common

import (
	"errors"
	"fmt"
	"os"
)

// FileFlag implements flag.Value
type FileFlag struct {
	filename string
	Content  []byte
}

func (f *FileFlag) String() string {
	return f.filename
}

func (f *FileFlag) Set(filename string) error {
	if filename == "" {
		return errors.New("unable to set a blank filename")
	}
	f.filename = filename

	b, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("unable to read: %w", err)
	}

	f.Content = b
	return nil
}
