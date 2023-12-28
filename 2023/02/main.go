package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/kmullin/advent-of-code/2023/common"
)

var (
	gameRe = regexp.MustCompile(`^Game (\d+): (.+)$`)
	cubeRe = regexp.MustCompile(`(\d+) (red|green|blue)(;|,)?`)
)

// game implements TextUnmarshaler
type game struct {
	ID   int
	Sets []cubeSet
}

func (g *game) AddSet(set cubeSet) {
	g.Sets = append(g.Sets, set)
}

// IsPossible determines if the game was possible with the given counts of cubes
func (g *game) IsPossible(red, green, blue int) bool {
	for _, s := range g.Sets {
		if s.Red > red || s.Green > green || s.Blue > blue {
			return false
		}
	}
	return true
}

func (g *game) UnmarshalText(text []byte) error {
	// match the first outer grouping
	match := gameRe.FindSubmatch(text)
	if len(match) != 3 {
		return errors.New("regex does not match 3 captures")
	}

	// we get game ID from the first digit of the match
	n, err := strconv.Atoi(string(match[1]))
	if err != nil {
		return err
	}
	g.ID = n

	var set cubeSet
	// parse the rest of the line
	for _, match := range cubeRe.FindAllSubmatch(match[2], -1) {
		if len(match) != 4 {
			return errors.New("regex does not match 4 captures")
		}
		color := string(match[2])
		count, err := strconv.Atoi(string(match[1]))
		if err != nil {
			return err
		}

		switch color {
		case "red":
			set.Red = count
		case "green":
			set.Green = count
		case "blue":
			set.Blue = count
		default:
			return fmt.Errorf("unknown color %s", color)
		}

		switch string(match[3]) {
		case ";", "":
			// zero out our set once we reach the end
			g.AddSet(set)
			set = cubeSet{}
		}
	}

	return nil
}

// cubeSet holds counts for each cube in the set
type cubeSet struct {
	Green, Red, Blue int
}

func main() {
	var filename common.FileFlag
	flag.Var(&filename, "input-file", "what")
	flag.Parse()

	var games []game
	scanner := bufio.NewScanner(bytes.NewReader(filename.Content))
	for scanner.Scan() {
		var g game
		err := g.UnmarshalText(scanner.Bytes())
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}
		games = append(games, g)
	}

	fmt.Printf("found %d games\n", len(games))

	// games possible
	fmt.Printf("sum of games that were possible: %d\n", gamesPossible(games))
}

func gamesPossible(games []game) (sum int) {
	for _, g := range games {
		if g.IsPossible(12, 13, 14) {
			sum += g.ID
		}
	}
	return
}
