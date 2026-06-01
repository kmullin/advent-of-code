package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"

	"github.com/kmullin/advent-of-code/internal/cli"
	"github.com/rs/zerolog/log"
)

func ReadInput(r io.Reader) (left []int, right []int, err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	start := true
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, nil, err
		}

		switch start {
		case true:
			left = append(left, n)
			start = false
		case false:
			right = append(right, n)
			start = true
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	Sort(left)
	Sort(right)
	return
}

func Sort(n []int) {
	if !slices.IsSorted(n) {
		slices.Sort(n)
	}
}

func GetDistance(l, r []int) int {
	var distance float64
	for i := range len(l) {
		distance += math.Abs(float64(l[i]) - float64(r[i]))
	}

	return int(distance)
}

func count(n int, x []int) (count int) {
	for i := range len(x) {
		if x[i] == n {
			count++
		}
	}
	return
}

func GetSimilarity(l, r []int) (similarity int) {
	for i := range len(l) {
		// how many times does the left int appear in the right
		similarity += l[i] * count(l[i], r)
	}

	return
}

func part1(b []byte) (any, error) {
	l, r, err := ReadInput(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	distance := GetDistance(l, r)
	return fmt.Sprintf("Total Distance: %v", distance), nil
}

func part2(b []byte) (any, error) {
	l, r, err := ReadInput(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	similarity := GetSimilarity(l, r)
	return fmt.Sprintf("Similarity: %v\n", similarity), nil
}

func main() {
	if err := cli.NewCmd(2024, 1, part1, part2).Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
	}

}
