package main

import (
	"bytes"
	"flag"
	"fmt"
	"regexp"
	"strconv"

	"github.com/kmullin/advent-of-code/2023/common"
)

var digitRe = regexp.MustCompile(`\d+`)

// schematic represents the input puzzle
type schematic struct {
	rows      [][]byte        // the raw data
	symbolMap map[coord][]int // map of symbol coordinates to a slice of numbers found around it
}

// coord is a simple coordinate tuple
type coord struct {
	X, Y int
}

// UnmarshalText implements the encoding.TextUnmarshaler interface
func (s *schematic) UnmarshalText(text []byte) error {
	s.symbolMap = make(map[coord][]int)
	s.rows = bytes.Split(text, []byte("\n"))
	s.findSymbols()
	s.findEdges()

	return nil
}

func (s *schematic) findSymbols() {
	// for each row, we walk byte by byte to find anything that is not a number or a period
	for r, line := range s.rows {
		for c, b := range line {
			// if 0-9 or .
			if (b >= 48 && b <= 57) || b == 46 {
				continue
			}
			s.symbolMap[coord{r, c}] = make([]int, 0)
		}
	}
}

func (s *schematic) findEdges() {
	// now grab edges of all the numbers
	for r, line := range s.rows {
		for _, match := range digitRe.FindAllIndex(line, -1) {
			var edges []coord
			// generate all surrounding coordinates from the matched number
			for row := r - 1; row <= r+1; row++ {
				for col := match[0] - 1; col < match[1]+1; col++ {
					edges = append(edges, coord{row, col})
				}
			}

			// now check each edge to see if it is also a location of a symbol
			for _, edge := range edges {
				for symbol := range s.symbolMap {
					if edge == symbol {
						// XXX: we already matched a digit regex so we shouldnt need to check err here
						n, _ := strconv.Atoi(string(line[match[0]:match[1]]))
						s.symbolMap[symbol] = append(s.symbolMap[symbol], n)
					}
				}
			}
		}
	}
}

func (s *schematic) SumOfPartNumbers() (sum int) {
	for _, nums := range s.symbolMap {
		for _, n := range nums {
			sum += n
		}
	}
	return
}

func (s *schematic) GearRatios() (sum int) {
	// find any symbol that has only 2 numbers near it
	for _, nums := range s.symbolMap {
		if len(nums) != 2 {
			continue
		}
		sum += nums[0] * nums[1]
	}
	return
}

func main() {
	var filename common.FileFlag
	flag.Var(&filename, "input-file", "what")
	flag.Parse()

	var s schematic
	s.UnmarshalText(filename.Content)
	fmt.Printf("sum of part numbers: %d\n", s.SumOfPartNumbers())
	fmt.Printf("sum of gear ratios: %d\n", s.GearRatios())
}
