package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []byte(`
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`)

func readInput(tb testing.TB, input []byte) (r ConditionRecord) {
	tb.Helper()
	err := r.UnmarshalText(input)
	assert.NoError(tb, err)
	return
}

func TestExample(t *testing.T) {
	r := readInput(t, exampleInput)

	// individual cases for each row
	cases := []int{
		1, 4, 1, 1, 4, 10,
	}
	for n, expected := range cases {
		t.Run(fmt.Sprintf("row%v", n), func(t *testing.T) {
			assert.Equal(t, expected, r.r[n].TotalArrangements())
		})
	}

	t.Run("total", func(t *testing.T) {
		assert.Equal(t, 21, r.TotalArrangements())
	})
}
