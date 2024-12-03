package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Reports []Report

// Report is a slice of ints, each int represents a Level
type Report []int

// IsSafe determines if the report is safe
func (r Report) IsSafe() bool {
	var isIncreasing bool

	for i, _ := range r {
		if i == len(r)-1 {
			continue
		}

		cur, next := float64(r[i]), float64(r[i+1])
		diff := cur - next
		abs := math.Abs(diff)

		// two adjacent levels differ by at least one and at most three
		if abs < 1 || abs > 3 {
			return false
		}

		// set first pass for increasing
		if i == 0 {
			isIncreasing = diff < 0
			continue
		}

		// all increasing or all decreasing
		// immediately bail if its different
		if (isIncreasing && diff > 0) || (!isIncreasing && diff < 0) {
			return false
		}
	}

	return true
}

func (r Reports) NumSafe() (count int) {
	for _, report := range r {
		if report.IsSafe() {
			count++
		}
	}

	return
}

func ReadInput(r io.Reader) (Reports, error) {
	var reports Reports
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		var report Report

		nums := strings.Split(scanner.Text(), " ")
		for _, num := range nums {
			i, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}

			report = append(report, i)
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func main() {
	log.SetFlags(0)

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reports, err := ReadInput(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of safe reports: %v", reports.NumSafe())
}
