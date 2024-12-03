package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const exampleInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func TestSafeReportsExample(t *testing.T) {
	reports, err := ReadInput(strings.NewReader(exampleInput))
	require.NoError(t, err)
	assert.Equal(t, 2, reports.NumSafe())
}
