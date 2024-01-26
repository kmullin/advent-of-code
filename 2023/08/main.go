package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/kmullin/advent-of-code/2023/common"
)

var pathRe = regexp.MustCompile(`^([\w]{3}) = \(([\w]{3}), ([\w]{3})\)$`)

type GhostMap struct {
	Instructions string // Instructions are our L/R moves
	Nodes        Network
}

type Network map[string]Path

type Path struct {
	Left, Right string
}

func (gm *GhostMap) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return errors.New("no content")
	}

	var line int
	gm.Nodes = make(Network)
	scanner := bufio.NewScanner(bytes.NewReader(text))
	for scanner.Scan() {
		line++
		if line == 1 {
			gm.Instructions = scanner.Text()
			continue
		}

		if match := pathRe.FindAllStringSubmatch(scanner.Text(), -1); match != nil {
			if len(match[0]) != 4 {
				return errors.New("regexp match doesnt have 3 groups")
			}
			gm.Nodes[match[0][1]] = Path{match[0][2], match[0][3]}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (gm *GhostMap) TotalSteps() int {
	return gm.findPath("AAA", func(s string) bool {
		if s == "ZZZ" {
			return true
		}
		return false
	})
}

func (gm *GhostMap) TotalStepsPart2() int {
	var startingPos []string
	for s, _ := range gm.Nodes {
		if strings.HasSuffix(s, "A") {
			startingPos = append(startingPos, s)
		}
	}

	var stepsToEnd []int
	for _, current := range startingPos {
		steps := gm.findPath(current, func(s string) bool {
			if strings.HasSuffix(s, "Z") {
				return true
			}
			return false
		})

		stepsToEnd = append(stepsToEnd, steps)
	}

	if len(stepsToEnd) >= 2 {
		return lcm(stepsToEnd[0], stepsToEnd[1], stepsToEnd[2:]...)
	}
	return stepsToEnd[0]
}

// findPath calculates how many steps it takes to get to something that makes endFunc return true
func (gm *GhostMap) findPath(start string, endFunc func(string) bool) (steps int) {
	current := start
	for {
		for _, r := range gm.Instructions {
			steps++
			switch r {
			case 'L':
				current = gm.Nodes[current].Left
			case 'R':
				current = gm.Nodes[current].Right
			default:
				panic(fmt.Sprintf("unknown step %q", r))
			}

			if endFunc(current) {
				return
			}
		}
	}
}

// gcd is greatest common denominator
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// lcm is lowest common multiple
func lcm(a, b int, integers ...int) int {
	result := a / gcd(a, b) * b

	for _, i := range integers {
		result = lcm(result, i)
	}

	return result
}

func main() {
	var filename common.FileFlag
	flag.Var(&filename, "input-file", "")
	flag.Parse()

	var gm GhostMap
	err := gm.UnmarshalText(filename.Content)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Number of steps: %v\n", gm.TotalSteps())
	fmt.Printf("Number of steps Part 2: %v\n", gm.TotalStepsPart2())
}
