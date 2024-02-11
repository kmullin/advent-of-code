package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kmullin/advent-of-code/2023/common"
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
