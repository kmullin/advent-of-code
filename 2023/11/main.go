package main

import (
	"bytes"
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/kmullin/advent-of-code/2023/common"
)

type Image struct {
	i                    [][]byte
	Galaxies             map[int]coord // map of galaxy number to coords
	spaceRows, spaceCols []int         // indices rows and cols with all dots
}

type coord struct {
	Row, Col int
}

func (i *Image) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return common.InputEmptyErr
	}
	for _, b := range bytes.Split(text, []byte("\n")) {
		if len(b) == 0 {
			continue
		}
		i.i = append(i.i, b)
	}
	i.Galaxies = make(map[int]coord)
	i.spaceRows, i.spaceCols = i.findEmptySpace()
	return nil
}

// findGalaxies finds all the coordinates of the galaxies after accounting for expanded space
// expansionFactor is the factor of how much space expands (part 2)
// part 1 expands by 2x
// part 2 expands by 1_000_000x
func (i *Image) findGalaxies(expansionFactor int) {
	var num int
	for row := 0; row < len(i.i); row++ {
		rowFactor := between(row, i.spaceRows)
		for col := 0; col < len(i.i[0]); col++ {
			colFactor := between(col, i.spaceCols)
			if i.i[row][col] == '#' {
				num++
				actualRow := row + rowFactor*(expansionFactor-1)
				actualCol := col + colFactor*(expansionFactor-1)
				i.Galaxies[num] = coord{actualRow, actualCol}
			}
		}
	}
}

// between calculates how many ints in the given indexes are between 0 and n
func between(n int, indexes []int) (count int) {
	for _, i := range indexes {
		if i <= n {
			count++
		}
	}
	return
}

// findEmptySpace returns the index of the rows and index of the columns where there are only dots
func (i Image) findEmptySpace() ([]int, []int) {
	rows := walk(len(i.i), len(i.i[0]), func(row, col int) bool {
		if i.i[row][col] != '.' {
			return false
		}
		return true
	})
	cols := walk(len(i.i[0]), len(i.i), func(col, row int) bool {
		if i.i[row][col] != '.' {
			return false
		}
		return true
	})

	return rows, cols
}

// ShortestPathSum calculates the sum of the paths between all the pairs of galaxies
// it loops thru all possible non-repeating pairs of galaxies and sums their path distances
func (i Image) ShortestPathSum() (sum int) {
	for a := 1; a <= len(i.Galaxies); a++ {
		for b := len(i.Galaxies); b > a; b-- {
			sum += i.shortestPath(a, b)
		}
	}
	return
}

// shortestPath determines the shortest path between 2 galaxies
// a is the start, b is the destination
// uses Taxicab Geometry: https://en.wikipedia.org/wiki/Taxicab_geometry
func (i Image) shortestPath(a, b int) int {
	x, y := i.Galaxies[a], i.Galaxies[b]
	return Abs(x.Row-y.Row) + Abs(x.Col-y.Col)
}

// walk is an attempt at being generic enough to loop thru a 2d array in both directions
// a and b are max values of the nested loop
func walk(a, b int, hasAllDots func(int, int) bool) (ii []int) {
	for i := 0; i < a; i++ {
		allDots := true
		for j := 0; j < b; j++ {
			if !hasAllDots(i, j) {
				allDots = false
				break
			}
		}
		if allDots {
			// we keep track of the outer loop iterator
			ii = append(ii, i)
		}
	}
	return
}

// pairCombinations calculates C(n,r)
func pairCombinations(n int) int {
	r := 2.0 // we have pairs
	ans := 1.0
	for i := 1.0; i <= r; i++ {
		ans *= float64(n) - r + i
		ans /= i
	}
	return int(ans)
}

// absolute value but with ints
func Abs(a int) int {
	return int(math.Abs(float64(a)))
}

func main() {
	var filename common.FileFlag
	flag.Var(&filename, "input-file", "")
	flag.Parse()

	var image Image
	err := image.UnmarshalText(filename.Content)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	for part, expansion := range []int{2, 1_000_000} {
		image.findGalaxies(expansion)
		fmt.Printf("Part %v, Factor: %v\n", part+1, expansion)
		fmt.Printf("Number of galaxies: %v\n", len(image.Galaxies))
		fmt.Printf("Number of pairs: %v\n", pairCombinations(len(image.Galaxies)))
		fmt.Printf("Shortest path sum: %v\n", image.ShortestPathSum())
		fmt.Println("")
	}
}
