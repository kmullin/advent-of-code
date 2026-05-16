package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleText = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

func TestExample(t *testing.T) {
	ranges := ReadInput(exampleText)

	assert.Equal(t, 1227775554, AddInvalidIDs(ranges, part1))
}

func TestExampleP2(t *testing.T) {
	ranges := ReadInput(exampleText)

	assert.Equal(t, 4174379265, AddInvalidIDs(ranges, part2))
}
