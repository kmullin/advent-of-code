package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput1 = []byte(`.....
.S-7.
.|.|.
.L-J.
.....`)

var exampleInput2 = []byte(`..F7.
.FJ|.
SJ.L7
|F--J
LJ...`)

var examplePart2Input1 = []byte(`...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`)

var examplePart2Input2 = []byte(`.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`)

var examplePart2Input3 = []byte(`FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`)

func readInput(t testing.TB, input []byte) (tiles Tiles) {
	t.Helper()
	err := tiles.UnmarshalText(input)
	assert.NoError(t, err)
	return
}

func TestFurthestPoint(t *testing.T) {
	cases := []struct {
		Input    []byte
		Expected int
	}{
		{exampleInput1, 4},
		{exampleInput2, 8},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			tiles := readInput(t, tc.Input)
			assert.Len(t, tiles, 5)
			steps, err := tiles.FurthestPoint()
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, steps)
		})
	}
}

func TestFindStart(t *testing.T) {
	cases := []struct {
		Input    []byte
		Expected coord
	}{
		{exampleInput1, coord{1, 1}},
		{exampleInput2, coord{2, 0}},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			tiles := readInput(t, tc.Input)
			s, err := tiles.findStart()
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, s)
		})
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		Input    []byte
		Expected int
	}{
		{examplePart2Input1, 4},
		{examplePart2Input2, 8},
		{examplePart2Input3, 10},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			tiles := readInput(t, tc.Input)
			s, err := tiles.InsideTiles()
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, s)
		})
	}
}

func TestPicksTheorem(t *testing.T) {
	cases := []struct {
		A, B int // area, boundary points
		I    int // interior points
	}{
		{10, 8, 7}, // examples on the wikipedia page
		{48, 96, 1},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.I, picksTheorem(tc.A, tc.B))
	}
}

func TestShoelace(t *testing.T) {
	// example coords taken from a youtube video I had to look up because
	// I have no idea how to math
	// https://www.youtube.com/watch?v=iKIpraBC-Nw
	cases := []struct {
		C        []coord
		Expected int
	}{
		{[]coord{{1, 6}, {3, 1}, {7, 2}, {4, 4}, {8, 5}}, 16}, // wikipedia
		{[]coord{{1, 6}, {-4, 3}, {-5, -3}, {3, -1}}, 43},     // youtube
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.Expected, shoelaceArea(tc.C))
		})
	}
}

func BenchmarkWalkPath(b *testing.B) {
	cases := [][]byte{
		exampleInput1,
		exampleInput2,
		examplePart2Input1,
		examplePart2Input2,
		examplePart2Input3,
	}

	for _, tc := range cases {
		b.Run("", func(b *testing.B) {
			tiles := readInput(b, tc)
			start, err := tiles.findStart()
			if err != nil {
				b.Fatal()
			}
			for i := 0; i < b.N; i++ {
				tiles.walkPath(start)
			}
		})
	}
}
