package main

import (
	"log"
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
	log.SetFlags(0)
	reports, err := ReadInput(strings.NewReader(exampleInput))
	require.NoError(t, err)

	t.Run("part 1", func(t *testing.T) {
		assert.Equal(t, 2, reports.NumSafe(1))
	})

	t.Run("part 2", func(t *testing.T) {
		assert.Equal(t, 4, reports.NumSafe(2))
	})
}
