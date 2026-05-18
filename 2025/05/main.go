package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"

	"github.com/kmullin/advent-of-code/internal/cli"
	"github.com/kmullin/advent-of-code/internal/common"
)

type Range struct {
	Start, End int
}

func (r Range) IsIn(n int) bool {
	return n >= r.Start && n <= r.End
}

func (r Range) Count() int {
	return r.End - r.Start + 1
}

func scanRanges(scanner *bufio.Scanner) ([]Range, error) {
	var rg []Range
	for scanner.Scan() {
		text := scanner.Text()

		// line break
		if len(text) == 0 {
			break
		}

		s := strings.Split(text, "-")
		if len(s) != 2 {
			return nil, fmt.Errorf("weird range: %q", string(text))
		}

		start, end := common.MustAtoi(s[0]), common.MustAtoi(s[1])
		rg = append(rg, Range{start, end})
	}

	return rg, nil
}

func findAllIngredients(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	ranges, err := scanRanges(scanner)
	if err != nil {
		return 0, err
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scanner err: %w", err)
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	total := 0
	curStart, curEnd := ranges[0].Start, ranges[0].End
	for _, r := range ranges {
		// are we overlapping
		if r.Start <= curEnd+1 {
			if r.End > curEnd {
				curEnd = r.End
			}
			continue
		}

		total += curEnd - curStart + 1
		curStart = r.Start
		curEnd = r.End
	}

	// final
	total += curEnd - curStart + 1
	return total, nil
}

func findFreshIngredients(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	rg, err := scanRanges(scanner)
	if err != nil {
		return 0, err
	}

	var freshIngredients int
	for scanner.Scan() {
		n := common.MustAtoi(scanner.Text())
		for _, r := range rg {
			if r.IsIn(n) {
				freshIngredients++
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scanner err: %w", err)
	}

	return freshIngredients, nil
}

func main() {
	ctx, err := cli.Setup(nil)
	if err != nil {
		log.Fatal(err)
	}

	freshIngredients, err := findFreshIngredients(ctx.Reader())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %v\n", freshIngredients)

	allIngredients, err := findAllIngredients(ctx.Reader())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("part 2: %v\n", allIngredients)
}
