package main

import (
	"strings"
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
	n := ReadInput(strings.NewReader(exampleInput))
	assert.Equal(t, 1, n)
}
