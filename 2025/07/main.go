package main

import (
	"bytes"

	"github.com/kmullin/advent-of-code/internal/cli"
	"github.com/rs/zerolog/log"
)

const (
	splitterCell = '^'
	rayCell      = '|'
	startCell    = 'S'
)

func part1(b []byte) (any, error) {
	var numSplits int

	lines := bytes.Split(bytes.TrimSpace(b), []byte("\n"))
	log.Debug().Int("total rows", len(lines)).Int("total cols", len(lines[0])).Msg("")
	for row, line := range lines {
		// we dont care about the second to last row
		if row >= len(lines)-1 {
			break
		}

		// found our starting point
		if i := bytes.Index(line, []byte{startCell}); i != -1 {
			lines[row+1][i] = rayCell // set the first ray
			continue
		}

		for col, b := range line {
			// if we arent a ray, skip
			if b != rayCell {
				continue
			}

			// check next row for splitter
			if lines[row+1][col] == splitterCell {
				log.Debug().Int("col", col).Int("row", row).Msg("found ray split")
				numSplits += 1
				lines[row+1][col-1] = rayCell
				lines[row+1][col+1] = rayCell
			} else {
				log.Debug().Int("col", col).Int("row", row).Msg("found ray continue")
				lines[row+1][col] = rayCell
			}
		}
	}

	return numSplits, nil
}

func depthFirstSearch() {

}

func part2(b []byte) (any, error) {
	return nil, nil
}

func main() {
	cmd := cli.NewCmd(2025, 7, part1, part2)
	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
