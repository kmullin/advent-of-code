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
type Report struct {
	Levels []int
}

// Safe determines if the report is safe, returns 0 if there were no failed checks, otherwise >0 will be a count of how many failed safety checks there were.
func (r *Report) Safe(part int) bool {
	var isIncreasing bool

	for i := range r.Levels {
		if i == len(r.Levels)-1 {
			continue
		}

		cur, next := float64(r.Levels[i]), float64(r.Levels[i+1])
		diff := cur - next
		abs := math.Abs(diff)

		// set first pass for increasing
		if i == 0 {
			isIncreasing = diff < 0
		}

		log.Printf("part: %v cur: %v, next: %v, increasing: %v",
			part, cur, next, isIncreasing,
		)

		// two adjacent levels differ by at least one and at most three
		if abs < 1 || abs > 3 {
			if part == 1 {
				return false
			}
			continue
		}

		// all increasing or all decreasing
		// immediately bail if its different
		if (isIncreasing && diff > 0) || (!isIncreasing && diff < 0) {
			if part == 1 {
				return false
			}
			continue
		}
	}

	return true
}

func (r Reports) NumSafe(part int) (count int) {
	for _, report := range r {
		n := report.Safe(part)
		if n {
			count++
		}
		log.Printf("%v: safe %v", report, n)
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

			report.Levels = append(report.Levels, i)
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

	log.Printf("Number of safe reports: %v", reports.NumSafe(1))
	log.Printf("Number of safe reports (Part 2): %v", reports.NumSafe(2))
}
