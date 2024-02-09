package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kmullin/advent-of-code/2023/common"
)

// various spring representations
const (
	operationalSpring = 0x2e // '.'
	damagedSpring     = 0x23 // '#'
	unknownSpring     = 0x3f // '?'
)

func main() {
	var filename common.FileFlag
	flag.Var(&filename, "input-file", "")
	flag.Parse()

	var record ConditionRecord
	err := record.UnmarshalText(filename.Content)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total Arrangements: %v\n", record.TotalArrangements())
}
