package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kmullin/advent-of-code/internal/cli"
	"github.com/kmullin/advent-of-code/internal/common"
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
			log.Fatalf("unexpected split length: %q", r)
		}

		start, end := common.MustAtoi(s[0]), common.MustAtoi(s[1])
		ranges = append(ranges, Range{start, end})
	}

	return ranges
}

// works by simply comparing the first half to the second half
// its naive and does not work for part 2
func part1(i int) bool {
	s := strconv.Itoa(i)
	n := len(s)

	// A repeated sequence twice must have even length.
	if n%2 != 0 {
		return false
	}

	half := n / 2
	if s[:half] == s[half:] {
		return true
	}

	return false
}

func part2(i int) bool {
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

func main() {
	ctx, err := cli.Setup(nil)
	if err != nil {
		log.Fatal(err)
	}

	ranges := ReadInput(string(ctx.Bytes()))

	fmt.Printf("Part 1: %v\n", AddInvalidIDs(ranges, part1))
	fmt.Printf("Part 2: %+v\n", AddInvalidIDs(ranges, part2))
}
