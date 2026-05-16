package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/kmullin/advent-of-code/internal/cli"
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
		log.Fatal(err)
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
		log.Fatal(err)
	}

	sum := 0
	for _, j := range joltages {
		sum += j
	}

	return sum
}

func main() {
	ctx, err := cli.Setup(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("part1 joltages: %v\n", ReadInput(ctx.Reader(), 2))
	fmt.Printf("part2 joltages: %v\n", ReadInput(ctx.Reader(), 12))
}
