package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleInput = `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   + `

func TestExample(t *testing.T) {
	total := ReadInput([]byte(exampleInput))
	assert.Equal(t, 4277556, total)
}

func TestExampleP2(t *testing.T) {
	total := readRightToLeft([]byte(exampleInput))
	assert.Equal(t, 3263827, total)
}
