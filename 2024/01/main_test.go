package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const exampleInput = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestExampleInput(t *testing.T) {
	r := strings.NewReader(exampleInput)
	left, right, err := ReadInput(r)

	require.NoError(t, err)

	t.Logf("%+v", left)
	t.Logf("%+v", right)
	i := GetDistance(left, right)
	assert.Equal(t, 11, i)
}
