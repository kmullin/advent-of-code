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
func (r *Report) Safe() int {
	var incidentCount int
	var isIncreasing bool

	for i, _ := range r.Levels {
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

		log.Printf("cur: %v, next: %v,increasing: %v", cur, next, isIncreasing)

		// two adjacent levels differ by at least one and at most three
		if abs < 1 || abs > 3 {
			log.Printf("inc abs: %v", abs)
			incidentCount++
			continue
		}

		// all increasing or all decreasing
		// immediately bail if its different
		if (isIncreasing && diff > 0) || (!isIncreasing && diff < 0) {
			log.Printf("inc diff: %v", diff)
			incidentCount++
			continue
		}
	}

	return incidentCount
}

func (r Reports) NumSafe(part int) (count int) {
	for _, report := range r {
		safe := "unsafe"

		n := report.Safe()
		switch {
		case n == 0:
			safe = "safe"
			count++
		case part == 2 && n > 1:
			log.Printf("incidentCount: %v", n)
		}
		log.Printf("%v: %v (%v)", report, safe, n)
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
