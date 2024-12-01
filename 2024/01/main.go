package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
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

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	l, r, err := ReadInput(f)
	if err != nil {
		log.Fatal(err)
	}

	distance := GetDistance(l, r)
	fmt.Printf("Total Distance: %v\n", distance)
}
