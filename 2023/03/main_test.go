package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleInput = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestExampleSchematic(t *testing.T) {
	var s schematic
	s.UnmarshalText([]byte(exampleInput))
	assert.Equal(t, 4361, s.SumOfPartNumbers(), "sum of part numbers incorrect")
}
