package common

import (
	"bytes"
	"fmt"
	"iter"
)

type Grid struct {
	cells [][]byte
}

type Coord struct {
	X, Y int
}

type Cell struct {
	Coord
	Val byte
}

func (c Coord) String() string {
	return fmt.Sprintf("(%v,%v)", c.X, c.Y)
}

func (c Cell) String() string {
	return fmt.Sprintf("(%v, %v): %v", c.X, c.Y, c.Val)
}

func NewGrid(b []byte) (*Grid, error) {
	g := &Grid{bytes.Split(b, []byte("\n"))}
	if !g.validate() {
		return nil, fmt.Errorf("matrix not the correct shape")
	}

	return g, nil
}

// validate ensures that the matrix is a rectangle
func (g *Grid) validate() bool {
	h := len(g.cells)

	if h < 2 {
		return false
	}

	w := len(g.cells[0])
	for _, row := range g.cells[min(1, len(g.cells)):] {
		if len(row) != w {
			return false
		}
	}

	return true
}

// Edges returns the edges of c. c can be off grid coordinates
func (g *Grid) Edges(c Coord, diagonal bool) []Coord {
	// our static 4 adjacent coords
	edges := []Coord{
		{c.X - 1, c.Y}, // up
		{c.X, c.Y - 1}, // left
		{c.X + 1, c.Y}, // down
		{c.X, c.Y + 1}, // right
	}

	if diagonal {
		corners := []Coord{
			{c.X - 1, c.Y - 1}, // up left
			{c.X - 1, c.Y + 1}, // up right
			{c.X + 1, c.Y - 1}, // down left
			{c.X + 1, c.Y + 1}, // down right
		}
		edges = append(edges, corners...)
	}

	numRows, numCols := len(g.cells)-1, len(g.cells[0])-1
	n := 0
	for _, edge := range edges {
		if edge.X < 0 || edge.X > numCols ||
			edge.Y < 0 || edge.Y > numRows {
			continue
		}

		edges[n] = edge
		n++
	}

	return edges[:n]
}

func (g *Grid) Get(c Coord) byte {
	return g.cells[c.Y][c.X]
}

func (g *Grid) Set(c Coord, b byte) {
	g.cells[c.Y][c.X] = b
}

func (g *Grid) Seq2() iter.Seq2[Coord, byte] {
	return func(yield func(Coord, byte) bool) {
		for y, row := range g.cells {
			for x, b := range row {
				if !yield(Coord{x, y}, b) {
					return
				}
			}
		}
	}
}

func (g *Grid) Seq() iter.Seq[Cell] {
	return func(yield func(Cell) bool) {
		for y, row := range g.cells {
			for x, b := range row {
				if !yield(Cell{
					Coord: Coord{
						X: x,
						Y: y,
					},
					Val: b,
				}) {
					return
				}
			}
		}
	}
}

func (g *Grid) String() string {
	return string(bytes.Join(g.cells, []byte("\n")))
}
