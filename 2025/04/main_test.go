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
	n := ReadInput([]byte(exampleInput))
	assert.Equal(t, 13, n)
}
