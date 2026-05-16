package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []byte(`987654321111111
811111111111119
234234234234278
818181911112111`)

func TestExample(t *testing.T) {
	sum := ReadInput(bytes.NewReader(exampleInput), 2)
	assert.Equal(t, 357, sum)
}

func TestExampleP2(t *testing.T) {
	sum := ReadInput(bytes.NewReader(exampleInput), 12)
	assert.Equal(t, 3121910778619, sum)
}
