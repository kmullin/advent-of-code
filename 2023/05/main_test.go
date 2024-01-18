package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func ReadInput(t *testing.T) *Almanac {
	t.Helper()

	var a Almanac
	err := a.UnmarshalText([]byte(exampleInput))
	assert.Nil(t, err)
	return &a
}

func TestAlmanac(t *testing.T) {
	a := ReadInput(t)

	cases := []struct {
		MapName  string
		Input    int
		Expected int
	}{
		{"seed-to-soil", 79, 81},
		{"seed-to-soil", 14, 14},
		{"seed-to-soil", 55, 57},
		{"seed-to-soil", 13, 13},

		{"soil-to-fertilizer", 15, 0},
		{"soil-to-fertilizer", 52, 37},
		{"soil-to-fertilizer", 0, 39},

		{"fertilizer-to-water", 53, 49},
		{"fertilizer-to-water", 11, 0},
		{"fertilizer-to-water", 0, 42},
		{"fertilizer-to-water", 7, 57},

		{"water-to-light", 18, 88},
		{"water-to-light", 25, 18},

		{"light-to-temperature", 77, 45},
		{"light-to-temperature", 45, 81},
		{"light-to-temperature", 64, 68},

		{"temperature-to-humidity", 69, 0},
		{"temperature-to-humidity", 0, 1},

		{"humidity-to-location", 56, 60},
		{"humidity-to-location", 93, 56},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%s-%d", tc.MapName, i), func(t *testing.T) {
			assert.Equal(t, tc.Expected, a.Lookup(tc.MapName, tc.Input))
		})
	}
}

func TestLowestLocation(t *testing.T) {
	a := ReadInput(t)

	t.Run("part1", func(t *testing.T) {
		assert.Equal(t, 35, a.LowestLocation())
	})

	t.Run("part2", func(t *testing.T) {
		assert.Equal(t, 46, a.LowestLocationP2())
	})
}
