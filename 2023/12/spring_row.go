package main

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// various spring representations
const (
	operationalSpring = "."
	damagedSpring     = "#"
	unknownSpring     = "?"
)

// unfold is the number of times to unfold the records for Part 2
const unfold = 5

type springRow struct {
	data           string
	damagedSprings []int // the size of each contiguous group of damaged springs is listed in the order those groups appear in the row
}

func newSpringRow(b []byte) (sr springRow, err error) {
	ss := strings.Fields(string(b))
	if len(ss) != 2 {
		return sr, errors.New("fields on row != 2")
	}
	sr.data = ss[0]
	sr.damagedSprings, err = convertCsvInts(ss[1])
	return sr, err
}

func (sr springRow) TotalArrangements() (count int) {
	return sr.recurse(0, "")
}

func (sr springRow) nonRecurse() (count int) {
	stack := []tup{{0, ""}}
	for len(stack) > 0 {
		t := stack[len(stack)-1] // pop
		stack = stack[:len(stack)-1]

		if t.i == len(sr.data) {
			if sr.isValidCombo(t.s) {
				count++
			}
			continue
		}

		if s := string(sr.data[t.i]); s != unknownSpring {
			stack = append(stack, tup{t.i + 1, t.s + s})
		} else {
			// we need to check for possible # and .
			t1 := tup{t.i + 1, t.s + damagedSpring}
			t2 := tup{t.i + 1, t.s + operationalSpring}
			stack = append(stack, t1, t2)
		}
	}

	return
}

func (sr springRow) recurse(i int, s string) (n int) {
	// fmt.Printf("(%v, %v)\t%v\n", i, s, sr.damagedSprings)
	if i == len(sr.data) {
		if sr.isValidCombo(s) {
			n = 1
		} else {
			n = 0
		}
	} else {
		if ss := string(sr.data[i]); ss != unknownSpring {
			n = sr.recurse(i+1, s+ss)
		} else {
			n = sr.recurse(i+1, s+damagedSpring) + sr.recurse(i+1, s+operationalSpring)
		}
	}

	return n
}

func (sr springRow) isValidCombo(s string) bool {
	var i []int
	for _, ss := range strings.Split(s, operationalSpring) {
		if len(ss) == 0 {
			continue
		}
		i = append(i, len(ss))
	}
	return slices.Compare(i, sr.damagedSprings) == 0
}

func (sr springRow) Unfold(expand int) springRow {
	l := len(sr.damagedSprings)

	var b strings.Builder
	for i := 1; i <= expand; i++ {
		fmt.Fprint(&b, sr.data)
		if i < expand {
			// if we're not the last iteration
			fmt.Fprint(&b, unknownSpring)
			sr.damagedSprings = append(sr.damagedSprings, sr.damagedSprings[:l]...)
		}
	}
	sr.data = b.String()
	return sr
}

// String implements the Stringer interface and allows us to pretty print the row
func (sr springRow) String() string {
	return fmt.Sprint(sr.data, " ", sr.damagedSprings)
}

// tup is a simple tuple that we can use to keep track of index and string to check
// during our iterations
type tup struct {
	i int
	s string
}

func (t tup) String() string {
	return fmt.Sprintf("(%v, %q)", t.i, t.s)
}

// convertCsvInts converts a byte slice of comma separated integers and returns an int slice
func convertCsvInts(s string) (i []int, err error) {
	var n int
	for _, ss := range strings.Split(s, ",") {
		n, err = strconv.Atoi(ss)
		if err != nil {
			return
		}
		i = append(i, n)
	}
	return
}
