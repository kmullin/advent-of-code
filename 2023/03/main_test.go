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

func testSetup(t *testing.T) *schematic {
	t.Helper()
	var s schematic

	s.UnmarshalText([]byte(exampleInput))
	return &s
}

func TestExampleSchematic(t *testing.T) {
	s := testSetup(t)
	assert.Equal(t, 4361, s.SumOfPartNumbers(), "sum of part numbers incorrect")
}

func TestExampleSchematicGearRatios(t *testing.T) {
	s := testSetup(t)
	assert.Equal(t, 467835, s.GearRatios(), "sum of all gear ratios")
}
