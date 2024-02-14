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
	unfold         bool  // to unfold or not
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
	return recurse(sr.data, sr.damagedSprings)
}

// nonRecurse was the initial implementation using a stack and a new tuple type
// it also cant easily be cached and its working off indexes which is useless for caching
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

// recurse was the second attempt of solving the problem, it cant easily be cached
// since its working on internal state of sr.data and sr.damagedSprings
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

// recurse is the third attempt at making a more cacheable implementation
// this one is passed the full string data, and works by walking left to right thru the characters while testing against the groups, only passing the remaining groups left as groups
// this should be more cacheable so we can solve part 2
func recurse(record string, groups []int) int {
	// check the cache here for a hit return instead of executing function
	f := func(record string, groups []int) int {
		// fmt.Printf("%v\t%v\n", record, groups)

		// we dont have any more groups that are desired
		// but we have springs in our record
		if len(groups) == 0 {
			// this catches if our record is empty
			if !strings.Contains(record, damagedSpring) {
				return 1
			}
			// we have damaged springs still, but no groups
			return 0
		}

		// no more springs, cannot proceed
		if len(record) == 0 {
			return 0
		}

		// grab our next work units
		c := string(record[0])
		g := groups[0]

		// inline function that uses given record and groups
		pound := func() int {
			// if we have a damaged spring, then the first group
			// should be able to be treated as a damaged spring
			var thisGroup string
			if len(record) < g {
				thisGroup = replaceUnknowns(record[:len(record)])
			} else {
				thisGroup = replaceUnknowns(record[:g])
			}

			// do the test to see if thisGroup matches all damaged springs of
			// the length of our group
			if thisGroup != strings.Repeat(damagedSpring, g) {
				return 0
			}

			// if our entire record is the length of our group
			// then we're done
			if len(record) == g {
				if len(groups) == 1 {
					return 1
				}
				return 0
			}

			// check the next characters can be separators
			if strings.Contains(string(record[g]), operationalSpring) ||
				strings.Contains(string(record[g]), unknownSpring) {
				return recurse(record[g+1:], groups[1:])
			}

			return 0
		}

		switch c {
		case damagedSpring: // #
			return pound()
		case unknownSpring: // ?
			return recurse(record[1:], groups) + pound()
		case operationalSpring: // .
			return recurse(record[1:], groups)
		default:
			panic("unknown spring type")
		}

		return 0
	}(record, groups)

	return f
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

func replaceUnknowns(s string) string {
	return strings.Replace(s, unknownSpring, damagedSpring, -1)
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
