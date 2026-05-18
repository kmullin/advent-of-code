package main

import (
	"fmt"
	"log"

	"github.com/kmullin/advent-of-code/internal/cli"
	"github.com/kmullin/advent-of-code/internal/common"
)

const paperRoll = '@'

// findPaperRolls finds the rolls of paper which have fewer than 4 rolls of paper in the 8 adjacent positions
func findPaperRolls(b []byte) (int, error) {
	grid, err := common.NewGrid(b)
	if err != nil {
		return 0, fmt.Errorf("matrix is not a rectangle: %w", err)
	}

	return countSurroundingPaperRolls(grid, 3), nil
}

func countSurroundingPaperRolls(g *common.Grid, maxNum int) (totalCount int) {
	for c, b := range g.Seq2() {
		if b != paperRoll {
			continue
		}

		edges := g.Edges(c, true)
		matching := 0
		for _, edge := range edges {
			if g.Get(edge) == paperRoll {
				matching++
			}
		}

		if matching <= maxNum {
			totalCount++
		}
	}

	return
}

// findPaperRolls2 finds the rolls of paper which have fewer than 4 rolls of paper in the 8 adjacent positions repeated (of course)
func findPaperRolls2(b []byte) (int, error) {
	grid, err := common.NewGrid(b)
	if err != nil {
		return 0, fmt.Errorf("matrix is not a rectangle: %w", err)
	}

	totalCount := 0
	var removed []common.Coord
	for len(removed) != 0 || totalCount == 0 {
		removed = make([]common.Coord, 0) // reset our list
		for c, b := range grid.Seq2() {
			if b != paperRoll {
				continue
			}

			edges := grid.Edges(c, true)

			// count matching edges
			matching := 0
			for _, edge := range edges {
				if grid.Get(edge) == paperRoll {
					matching++
				}
			}

			if matching < 4 {
				removed = append(removed, c)
			}

		}

		// after walking the whole grid, we need to remove the paper rolls with a forklift
		for _, c := range removed {
			grid.Set(c, 'x')
		}
		totalCount += len(removed)
	}

	return totalCount, nil
}

func main() {
	ctx, err := cli.Setup(nil)
	if err != nil {
		log.Fatal(err)
	}

	count, err := findPaperRolls(ctx.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %v\n", count)

	count, err = findPaperRolls2(ctx.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 2: %v\n", count)
}
