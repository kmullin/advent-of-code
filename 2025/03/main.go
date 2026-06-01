package main

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"

	"github.com/kmullin/advent-of-code/internal/cli"
	"github.com/rs/zerolog/log"
)

type seen struct {
	num, pos int
}

// findMax returns the number of jolts made up from the bank of batteries enabled
func findMax(s string, nDigits int) int {
	var num []int
	lastSeen := seen{0, -1}
	for len(num) < nDigits {
		// create a dynamic window
		// we dont care about digits before the previous position
		start := max(0, lastSeen.pos+1)
		end := len(s) - nDigits + len(num)

		lastSeen = seen{0, -1}
		for i := start; i <= end; i++ {
			n := int(s[i] - '0') // subtract 48 from byte value

			if n > lastSeen.num {
				lastSeen.num, lastSeen.pos = n, i
			}
		}
		num = append(num, lastSeen.num)
	}

	// take slice of int and convert to string
	var sb strings.Builder
	for _, n := range num {
		sb.WriteString(strconv.Itoa(n))
	}

	// then back to int?
	n, err := strconv.Atoi(sb.String())
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	return n
}

func ReadInput(r io.Reader, nDigits int) int {
	scanner := bufio.NewScanner(r)

	var joltages []int
	for scanner.Scan() {
		joltages = append(joltages, findMax(scanner.Text(), nDigits))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal().Err(err).Msg("")
	}

	sum := 0
	for _, j := range joltages {
		sum += j
	}

	return sum
}

func part1(b []byte) (any, error) {
	return ReadInput(bytes.NewReader(b), 2), nil
}

func part2(b []byte) (any, error) {
	return ReadInput(bytes.NewReader(b), 12), nil
}

func main() {
	if err := cli.NewCmd(2025, 3, part1, part2).Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
