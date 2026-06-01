package main

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/kmullin/advent-of-code/internal/cli"
	"github.com/kmullin/advent-of-code/internal/common"
	"github.com/rs/zerolog/log"
)

func multiplyF(n, sum int) int {
	if sum == 0 {
		sum = 1
	}
	return sum * n
}

func sumF(n, sum int) int {
	return sum + n
}

func readInput(b []byte) ([][]int, []string) {
	s := strings.Split(string(b), "\n")

	// the last row are symbols, skip them
	grid := make([][]int, len(s)-1)
	for row, line := range s[:len(s)-1] {
		fields := strings.Fields(line)

		grid[row] = make([]int, len(fields))
		for col := range len(fields) {
			grid[row][col] = common.MustAtoi(fields[col])
		}
	}

	return grid, strings.Fields(s[len(s)-1])
}

func maxLineLen(bb [][]byte) int {
	m := 0
	for _, b := range bb {
		m = max(len(b), m)
	}
	return m
}

func readRightToLeft(b []byte) (any, error) {
	bb := bytes.Split(bytes.TrimRight(b, "\n"), []byte("\n"))
	symbols := bytes.Fields(bb[len(bb)-1])
	maxLen := maxLineLen(bb[:len(bb)-1])

	// we need to iterate backwards, but also whitespace might be missing from
	// the end of the row
	field := len(symbols) - 1
	total, grandTotal := 0, 0
	for col := maxLen - 1; col >= 0; col-- {
		var sb strings.Builder

		// now we start at the top and read downwards
		for row := 0; row <= len(bb)-2; row++ {
			if len(bb[row])-1 < col {
				// ignore missing whitespace at the end
				continue
			}

			log.Debug().
				Str("found", string(bb[row][col])).
				Int("row", row).
				Int("col", col).
				Msg("")
				// ("found:%q row:%v/%v col:%v/%v", bb[row][col], row, len(bb)-2, col, maxLen-1)

			if unicode.IsSpace(rune(bb[row][col])) {
				continue
			}

			sb.WriteByte(bb[row][col])
		}

		s := sb.String()
		if len(s) == 0 {
			// we've found the end of a field
			field--
			grandTotal += total
			total = 0
			continue
		}

		n := common.MustAtoi(s)
		switch r := symbols[field]; r[0] {
		case '*':
			total = multiplyF(n, total)
		case '+':
			total = sumF(n, total)
		}
		log.Debug().Int("total", total).Int("n", n).Int("col", col).Msg("")
	}

	return grandTotal + total, nil
}

func ReadInput(b []byte) (any, error) {
	things, symbols := readInput(bytes.TrimRight(b, "\n"))

	var mathFunc func(int, int) int
	grandTotal := 0
	// assumes same number of columns and rows
	for col, sym := range symbols {
		switch sym {
		case "*":
			mathFunc = multiplyF
		case "+":
			mathFunc = sumF
		}

		sum := 0
		for row := range len(things) {
			sum = mathFunc(things[row][col], sum)
		}
		grandTotal += sum

		//log.Printf("%v of col:%v %v", sym, col, sum)
	}

	return grandTotal, nil
}

func main() {
	cmd := cli.NewCmd(2025, 6, ReadInput, readRightToLeft)
	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
