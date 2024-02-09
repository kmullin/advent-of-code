package main

import "fmt"

// springTypes is a helpful wrapper for wording
var springTypes = map[byte]string{
	operationalSpring: "o",
	damagedSpring:     "d",
	unknownSpring:     "u",
}

// springGroup holds the indexes of the start and end of a grouping of springs
type springGroup struct {
	start, end int  // indexes of groupings
	springType byte // which type of grouping, instead of doing type assertions with subtypes
}

func newSpringGroup() springGroup {
	return springGroup{-1, -1, 0}
}

// Len returns the size of the range for the spring group
func (sg springGroup) Len() int {
	return sg.end - sg.start + 1
}

func (sg springGroup) String() string {
	return fmt.Sprintf("%v: %v (%v,%v)", springTypes[sg.springType], sg.Len(), sg.start, sg.end)
}
