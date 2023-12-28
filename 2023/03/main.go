package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"regexp"

	"github.com/kmullin/advent-of-code/2023/common"
)

var digitRe = regexp.MustCompile(`\d+`)

// schematic represents the input puzzle
type schematic struct {
	// the raw data
	rows [][]byte

	SymbolCoords []coord // coordinates of where all symbols are
	NumberCoords []coord // coordinates of where all the numbers are
}

type coord struct {
	X, Y int
}

// UnmarshalText implements the encoding.TextUnmarshaler interface
func (s *schematic) UnmarshalText(text []byte) error {
	scanner := bufio.NewScanner(bytes.NewReader(text))
	line := 0
	for scanner.Scan() {
		// for each row, we walk byte by byte to find anything that is not a number or a period
		row := scanner.Bytes()
		// we save the row as is for later
		s.rows = append(s.rows, row)
		for c, b := range row {
			// if 0-9 or .
			if (b >= 48 && b <= 57) || b == 46 {
				continue
			}
			s.SymbolCoords = append(s.SymbolCoords, coord{line, c})
		}

		// now grab edges of all the numbers
		for _, match := range digitRe.FindAllIndex(row, -1) {
			c := coord{line, match[0]}
			s.NumberCoords = append(s.NumberCoords, c)
			fmt.Printf("%v\n", c)
		}

		line++
	}

	return nil
}

func (s *schematic) GetAllPartNumbers() (partNums []int) {
	//for _, row := range s.rows {
	//}
	return
}

func (s *schematic) SumOfPartNumbers() (sum int) {
	for _, n := range s.GetAllPartNumbers() {
		sum += n
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
}
