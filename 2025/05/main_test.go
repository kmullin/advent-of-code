package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestExample(t *testing.T) {
	freshIngredients, err := findFreshIngredients(strings.NewReader(exampleInput))
	assert.NoError(t, err)
	assert.Equal(t, 3, freshIngredients)
}

func TestExampleP2(t *testing.T) {
	allIngredients, err := findAllIngredients(strings.NewReader(exampleInput))
	assert.NoError(t, err)
	assert.Equal(t, 14, allIngredients)
}
