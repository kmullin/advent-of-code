package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []byte(`
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`)

var exampleExpected = []byte(`
....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......`)

const partOneExpansionFactor = 2

const (
	partTwoExpansionFactor1 = 10
	partTwoExpansionFactor2 = 100
)

func readInput(tb testing.TB, input []byte) (i Image) {
	tb.Helper()
	err := i.UnmarshalText(input)
	assert.NoError(tb, err)
	return
}

func TestFindGalaxies(t *testing.T) {
	given := readInput(t, exampleInput)
	t.Run("Part 1", func(t *testing.T) {
		given.findGalaxies(partOneExpansionFactor)
		// known coords from our example (after expansion)
		var galaxyCoords = map[int]coord{
			1: {0, 4},
			2: {1, 9},
			3: {2, 0},
			4: {5, 8},
			5: {6, 1},
			6: {7, 12},
			7: {10, 9},
			8: {11, 0},
			9: {11, 5},
		}
		for num, c := range given.Galaxies {
			t.Log(c)
			assert.Equal(t, galaxyCoords[num], c)
		}
	})
	t.Run("Part 2 10", func(t *testing.T) {
	})
	t.Run("Part 2 100", func(t *testing.T) {
	})
}

func TestShortestPath(t *testing.T) {
	image := readInput(t, exampleInput)
	image.findGalaxies(partOneExpansionFactor)

	cases := []struct {
		A, B             int // the galaxy numbers
		ExpectedDistance int // the distance between A and B
	}{
		{1, 2, 6},
		{5, 9, 9},
		{1, 7, 15},
		{3, 6, 17},
		{8, 9, 5},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.ExpectedDistance, image.shortestPath(tc.A, tc.B))
		})
	}
}

func TestShortestPathSum(t *testing.T) {
	image := readInput(t, exampleInput)
	image.findGalaxies(partOneExpansionFactor)
	assert.Equal(t, 374, image.ShortestPathSum())
}

func TestPairCombinations(t *testing.T) {
	assert.Equal(t, 36, pairCombinations(9))
}

func TestBetween(t *testing.T) {
	cases := []struct {
		N        int
		Indexes  []int
		Expected int
	}{
		{10, []int{3, 4, 10}, 3},
		{10, []int{3, 4}, 2},
		{10, []int{3, 4, 11}, 2},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.Expected, between(tc.N, tc.Indexes))
		})
	}
}
