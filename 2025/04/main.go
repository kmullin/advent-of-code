package main

import (
	"bytes"
	"fmt"

	"github.com/kmullin/advent-of-code/internal/cli"
	"github.com/kmullin/advent-of-code/internal/common"
	"github.com/rs/zerolog/log"
)

const paperRoll = '@'

// findPaperRolls finds the rolls of paper which have fewer than 4 rolls of paper in the 8 adjacent positions
func findPaperRolls(b []byte) (any, error) {
	grid, err := common.NewGrid(bytes.TrimSpace(b))
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
func findPaperRolls2(b []byte) (any, error) {
	grid, err := common.NewGrid(bytes.TrimSpace(b))
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
	cmd := cli.NewCmd(2025, 4, findPaperRolls, findPaperRolls2)
	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
