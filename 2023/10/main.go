package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"slices"

	"github.com/kmullin/advent-of-code/2023/common"
)

// Tiles represents the tiles given as input
type Tiles [][]rune

// coord represents a point on a tile
type coord struct {
	Row, Col int
}

func (t *Tiles) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return common.InputEmptyErr
	}
	for _, b := range bytes.Split(text, []byte("\n")) {
		if len(b) == 0 {
			continue
		}
		*t = append(*t, []rune(string(b)))
	}

	return nil
}

// FurthestPoint determines the furthest point as a number of steps starting from S
func (t Tiles) FurthestPoint() (int, error) {
	start, err := t.findStart()
	if err != nil {
		return -1, fmt.Errorf("unable to find start: %w", err)
	}

	return len(t.walkPath(start)) / 2, nil
}

// InsideTiles calculates the number of tiles that are enclosed by the loop.
// It does this by first calculating the area using the shoestring formula.
// Then it use "Pick's Theorem" to solve for the number of points "inside".
// I had just learned about this due to a hint from reddit.
func (t Tiles) InsideTiles() (int, error) {
	start, err := t.findStart()
	if err != nil {
		return -1, fmt.Errorf("unable to find start: %w", err)
	}

	path := t.walkPath(start)
	return picksTheorem(shoelaceArea(path), len(path)), nil
}

// findStart returns the coord of where 'S' is on the tile
func (t Tiles) findStart() (coord, error) {
	for row, line := range t {
		for col, r := range line {
			if r == 'S' {
				return coord{row, col}, nil
			}
		}
	}
	return coord{-1, -1}, errors.New("unable to find 'S' in tiles")
}

// walkPath finds the edges starting from given coord, takes into account direction
func (t Tiles) walkPath(origin coord) (path []coord) {
	edges := t.findEdges(origin)

	path = append(path, origin) // record the whole path
	q := []coord{edges[0]}      // choose 1 edge for now, we know we're a loop
	for len(q) != 0 {
		edge := q[0] // popleft
		q = q[1:]    // readjust q

		// determine where we came from
		row, col := edge.Row-origin.Row, edge.Col-origin.Col
		switch r := t[edge.Row][edge.Col]; r {
		case 'S': // we might come back to this in a loop
			break
		case '|', '-':
			// | is a vertical pipe connecting north and south.
			// - is a horizontal pipe connecting east and west.
			q = append(q, coord{edge.Row + row, edge.Col + col})
		case 'L', '7':
			// L is a 90-degree bend connecting north and east.
			// 7 is a 90-degree bend connecting south and west.
			q = append(q, coord{edge.Row + col, edge.Col + row})
		case 'J', 'F':
			// J is a 90-degree bend connecting north and west.
			// F is a 90-degree bend connecting south and east.
			q = append(q, coord{edge.Row - col, edge.Col - row})
		default:
			panic("default case statement")
		}

		path = append(path, edge)
		origin = edge // keep track of our walking
	}

	return
}

func (t Tiles) findEdges(origin coord) (edges []coord) {
	// were are at the start we need to find all paths
	// we can only go 4 directions
	directions := map[string]coord{
		"north": {-1, 0},
		"south": {1, 0},
		"west":  {0, -1},
		"east":  {0, 1},
	}

	possible := map[string][]rune{
		"north": []rune{'|', '7', 'F'},
		"south": []rune{'|', 'J', 'L'},
		"west":  []rune{'-', 'F', 'L'},
		"east":  []rune{'-', 'J', '7'},
	}

	for direction, mod := range directions {
		c := coord{origin.Row + mod.Row, origin.Col + mod.Col}

		// bounds check
		if c.Row < 0 || c.Col < 0 || c.Row > len(t)-1 || c.Col > len(t[0])-1 {
			continue
		}

		// we need to check if the edges are valid
		if r := t[c.Row][c.Col]; r == '.' || !slices.Contains(possible[direction], r) {
			continue
		}

		edges = append(edges, c)
	}

	return
}

// picksTheorem is used to find the area of a polygon with integer vertex coordinates
// we have a non-simple polygon, but based on the problem description we shouldn't have any 'holes'
// returns number of points in the area
// https://en.wikipedia.org/wiki/Pick%27s_theorem
func picksTheorem(a, b int) int {
	// a = i + b/2 + h - 1
	// where:
	//   a = area
	//   b = boundary points
	//   i = points interior to the polygon
	//   h = number of holes
	h := 0                 // we have no holes
	return a - b/2 - h + 1 // returns i
}

// shoelaceArea implements the shoelace formula to calculate the area of a polygon
// https://en.wikipedia.org/wiki/Shoelace_formula
func shoelaceArea(coords []coord) (n int) {
	// we need the first element at the end too for this to work
	coords = append(coords, coords[0])
	for i := 0; i < len(coords)-1; i++ {
		n += coords[i].Row*coords[i+1].Col - coords[i].Col*coords[i+1].Row
	}

	return int(math.Abs(1.0 / 2.0 * float64(n)))
}

func main() {
	var filename common.FileFlag
	flag.Var(&filename, "input-file", "")
	flag.Parse()

	var tiles Tiles
	err := tiles.UnmarshalText(filename.Content)
	fatalErr(err)

	steps, err := tiles.FurthestPoint()
	fatalErr(err)
	fmt.Printf("Furthest Point: %v\n", steps)

	insideCount, err := tiles.InsideTiles()
	fatalErr(err)
	fmt.Printf("Inside Tiles: %v\n", insideCount)
}

func fatalErr(err error) {
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
}
