package cli

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spf13/pflag"
)

const InputFilename = "input.txt"

// Context is our CLI context, which holds the read input file
type Context struct {
	InputFilename string

	data []byte
}

// Setup sets up the CLI with default flags, optionally configures extra flags on the flagset, returns Context
func Setup(configure func(*pflag.FlagSet)) (*Context, error) {
	fs := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

	filename := fs.StringP("input-file", "f", InputFilename, "File with input data")

	if configure != nil {
		configure(fs)
	}

	if err := fs.Parse(os.Args[1:]); err != nil {
		return nil, fmt.Errorf("err parsing flags: %w", err)
	}

	b, err := os.ReadFile(*filename)
	if err != nil {
		return nil, fmt.Errorf("err reading input file: %w", err)
	}

	return &Context{
		InputFilename: *filename,
		data:          bytes.TrimSpace(b),
	}, nil
}

func (c *Context) Bytes() []byte {
	return c.data
}

func (c *Context) Reader() io.Reader {
	return bytes.NewReader(c.data)
}
