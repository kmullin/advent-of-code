package main

import (
	"bytes"
	"errors"
	"fmt"
	"slices"
	"strconv"
)

type springRow struct {
	data           []byte
	damagedSprings []int         // the size of each contiguous group of damaged springs is listed in the order those groups appear in the row
	groups         []springGroup // groupings of consecutive springs
}

func newSpringRow(b []byte) (sr springRow, err error) {
	bb := bytes.Fields(b)
	if len(bb) != 2 {
		return sr, errors.New("fields on row != 2")
	}
	sr.data = bb[0]
	sr.damagedSprings, err = convertCsvInts(bb[1])
	// here after because we best effort printing in String() err is not fatal
	sr.groups = sr.findGroupings()
	return sr, err
}

type tup struct {
	i int
	b []byte
}

func (t tup) String() string {
	return fmt.Sprintf("(%v, '%s')", t.i, t.b)
}

func (sr springRow) TotalArrangements() (count int) {
	var buf bytes.Buffer
	stack := []tup{{0, []byte{}}}
	for len(stack) > 0 {
		s := stack[len(stack)-1] // pop
		stack = stack[:len(stack)-1]

		if s.i == len(sr.data) {
			if isValid(s.b, sr.damagedSprings) {
				count++
			}
			continue
		}

		buf.Reset()           // fresh start
		_, _ = buf.Write(s.b) // write our current bytes to buffer
		if b := sr.data[s.i]; b != unknownSpring {
			_ = buf.WriteByte(b)
			stack = append(stack, tup{s.i + 1, buf.Bytes()})
		} else {
			b1, b2 := bytes.Clone(buf.Bytes()), bytes.Clone(buf.Bytes())
			// we need to check for possible # and .
			b1 = append(b1, damagedSpring)
			b2 = append(b2, operationalSpring)
			stack = append(stack, tup{s.i + 1, b1}, tup{s.i + 1, b2})
		}
	}

	return
}

func isValid(b []byte, groups []int) bool {
	var i []int
	for _, bb := range bytes.Split(b, []byte{operationalSpring}) {
		if len(bb) == 0 {
			continue
		}
		i = append(i, len(bb))
	}
	return slices.Compare(i, groups) == 0
}

func (sr springRow) filterGroups(f func(b byte) bool) (groups []springGroup) {
	for _, g := range sr.groups {
		if f(g.springType) {
			groups = append(groups, g)
		}
	}
	return
}

func (sr springRow) findGroupings() (groups []springGroup) {
	// so we cant run more than once
	if len(sr.groups) != 0 {
		return
	}

	sg := newSpringGroup()
	for i, b := range sr.data {
		if sg.start == -1 {
			sg.start = i
			sg.springType = b
		}
		sg.end = i

		// lookahead to see if we need to save our grouping
		if len(sr.data)-1 == i || sg.springType != sr.data[i+1] {
			groups = append(groups, sg)
			sg = newSpringGroup()
		}
	}
	return
}

// String implements the Stringer interface and allows us to pretty print the row
func (sr springRow) String() string {
	return fmt.Sprintf("%v %v (%v) %v", string(sr.data), sr.groups, len(sr.groups), sr.damagedSprings)
}

// convertCsvInts converts a byte slice of comma separated integers and returns an int slice
func convertCsvInts(b []byte) (i []int, err error) {
	var n int
	for _, bb := range bytes.Split(b, []byte(",")) {
		n, err = strconv.Atoi(string(bb))
		if err != nil {
			return
		}
		i = append(i, n)
	}
	return
}
