package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleInput = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestExample(t *testing.T) {
	n, err := findPaperRolls([]byte(exampleInput))
	assert.NoError(t, err)
	assert.Equal(t, 13, n)
}

func TestExampleP2(t *testing.T) {
	n, err := findPaperRolls2([]byte(exampleInput))
	assert.NoError(t, err)
	assert.Equal(t, 43, n)
}
