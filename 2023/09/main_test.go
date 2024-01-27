package main

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []byte(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`)

func TestExtrapolate(t *testing.T) {
	var o OASIS
	err := o.UnmarshalText(exampleInput)
	assert.NoError(t, err)

	assert.Equal(t, 114, o.Extrapolate(1))
	assert.Equal(t, 2, o.Extrapolate(2))
}

func TestFindSteps(t *testing.T) {
	var o OASIS
	err := o.UnmarshalText(exampleInput)
	assert.NoError(t, err)

	cases := []struct {
		Start    []int
		Expected []int
	}{
		{o.History[0], []int{0, 0, 0, 0}},
		{o.History[1], []int{0, 0, 0}},
		{o.History[2], []int{0, 0}},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("example%v", i+1), func(t *testing.T) {
			answer := tc.Start
			for slices.Max(answer) != 0 {
				answer = findSteps(answer)
			}
			assert.Equal(t, tc.Expected, answer)
		})
	}
}

func TestNegativeInput(t *testing.T) {
	cases := [][]int{
		{1, -2, 10, 50, 131, 266, 468, 750, 1125, 1606, 2206, 2938, 3815, 4850, 6056, 7446, 9033, 10830, 12850, 15106, 17611},
		{19, 40, 84, 160, 276, 435, 624, 790, 790, 291, -1419, -5767, -15459, -35380, -73969, -145243, -271683, -488238, -847750, -1428154, -2341862},
	}

	for _, nums := range cases {
		t.Run("line", func(t *testing.T) {
			// t.Log(nums)
			for slices.Max(nums) != 0 {
				nums = findSteps(nums)
				//t.Log(nums)
			}
		})
	}
}
