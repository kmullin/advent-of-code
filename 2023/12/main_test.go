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

func TestExamplePart1(t *testing.T) {
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

func TestExamplePart2(t *testing.T) {
	r := readInput(t, exampleInput)

	// individual cases for each row
	cases := []int{
		1, 16384, 1, 16, 2500, 506250,
	}
	for n, expected := range cases {
		t.Run(fmt.Sprintf("row%v", n), func(t *testing.T) {
			r.r[n].unfold = true // handle the unfolding of the spring rows
			assert.Equal(t, expected, r.r[n].TotalArrangements())
		})
	}

	t.Run("total", func(t *testing.T) {
		assert.Equal(t, 525152, r.TotalArrangements())
	})
}
