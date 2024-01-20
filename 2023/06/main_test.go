package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput = `Time:      7  15   30
Distance:  9  40  200`

func TestPartOne(t *testing.T) {
	var races Races
	err := races.UnmarshalText([]byte(exampleInput))
	assert.Nil(t, err)
	assert.Equal(t, 3, len(races))

	// first test to ensure our math is correct on a per race basis
	cases := []struct {
		Race             Race
		ExpectedWinnable int
	}{
		{races[0], 4},
		{races[1], 8},
		{races[2], 9},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("race-%d", i), func(t *testing.T) {
			assert.Equal(t, tc.ExpectedWinnable, tc.Race.WinnableTimings())
		})
	}
}

func TestPartTwo(t *testing.T) {
	var race Race
	err := race.UnmarshalText([]byte(exampleInput))
	assert.Nil(t, err)

	assert.Equal(t, 71503, race.WinnableTimings())
}
