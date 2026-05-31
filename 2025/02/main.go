package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/kmullin/advent-of-code/internal/cli"
	"github.com/kmullin/advent-of-code/internal/common"
	"github.com/rs/zerolog/log"
)

type Range struct {
	Start int
	End   int
}

func (r Range) String() string {
	return fmt.Sprintf("%v-%v", r.Start, r.End)
}

func ReadInput(input string) []Range {
	var ranges []Range
	for r := range strings.SplitSeq(input, ",") {
		s := strings.Split(r, "-")
		if len(s) != 2 {
			log.Fatal().Msgf("unexpected split length: %q", r)
		}

		start, end := common.MustAtoi(s[0]), common.MustAtoi(s[1])
		ranges = append(ranges, Range{start, end})
	}

	return ranges
}

// works by simply comparing the first half to the second half
// its naive and does not work for part 2
// Part 1
func repeatsTwice(i int) bool {
	s := strconv.Itoa(i)
	n := len(s)

	// A repeated sequence twice must have even length.
	if n%2 != 0 {
		return false
	}

	half := n / 2
	return s[:half] == s[half:]
}

func repeatsTwiceOrMore(i int) bool {
	s := strconv.Itoa(i)
	n := len(s)
	// Try every possible pattern length
	for k := 1; k <= n/2; k++ {

		// Pattern length must divide evenly
		if n%k != 0 {
			continue
		}

		pattern := s[:k]
		repeats := n / k

		// Rebuild the string
		if strings.Repeat(pattern, repeats) == s {
			return true
		}
	}

	return false
}

func AddInvalidIDs(ranges []Range, f func(int) bool) int {
	var ids []int
	for _, r := range ranges {
		// for every number in the range
		for i := r.Start; i <= r.End; i++ {
			if f(i) {
				ids = append(ids, i)
			}
		}
	}

	count := 0
	for _, n := range ids {
		count += n
	}
	return count
}

func part1(b []byte) (string, error) {
	ranges := ReadInput(string(bytes.TrimSpace(b)))
	return strconv.Itoa(AddInvalidIDs(ranges, repeatsTwice)), nil
}

func part2(b []byte) (string, error) {
	ranges := ReadInput(string(bytes.TrimSpace(b)))
	return strconv.Itoa(AddInvalidIDs(ranges, repeatsTwiceOrMore)), nil
}

func main() {
	if err := cli.NewCmd(2025, 2, part1, part2).Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
