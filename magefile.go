//go:build mage

//mage:multiline

// Advent of Code helpers
package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/kmullin/advent-of-code/internal/templates"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// New creates a new problem scaffold to get started (year, day)
func New(ctx context.Context, year, day int) error {
	path := filepath.Join(
		fmt.Sprintf("%d", year),
		fmt.Sprintf("%02d", day),
	)

	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("%v already exists", path)
	}

	fmt.Printf("Generating new problem for %v...\n", path)
	if err := os.MkdirAll(path, 0o755); err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}

	data := templates.Data{Year: year, Day: day}
	for _, file := range []string{"main.go", "main_test.go", "README.md"} {
		t := fmt.Sprintf("%v.tmpl", file)

		tmpl, err := template.ParseFS(templates.FS, t)
		if err != nil {
			return fmt.Errorf("unable to read from template %q: %w", t, err)
		}

		filename := filepath.Join(path, file)
		f, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("unable to create file: %w", err)
		}
		defer f.Close()

		fmt.Printf(" \\_ %v\n", filename)
		if err := tmpl.Execute(f, data); err != nil {
			return fmt.Errorf("unable to execute template: %w", err)
		}
	}

	return nil
}
