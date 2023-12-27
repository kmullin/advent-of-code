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
	Sets []Set
}

func (g *game) AddSet(set Set) {
	g.Sets = append(g.Sets, set)
}

func (g *game) UnmarshalText(text []byte) error {
	// match the first outer grouping
	match := gameRe.FindSubmatch(text)
	if len(match) != 3 {
		return errors.New("regex does not match 3 captures")
	}

	// we get game ID from the first digit of the match
	n, err := strconv.Atoi(fmt.Sprintf("%s", match[1]))
	if err != nil {
		return err
	}
	g.ID = n

	var set Set
	// parse the rest of the line
	for _, match := range cubeRe.FindAllSubmatch(match[2], -1) {
		if len(match) != 4 {
			return errors.New("regex does not match 4 captures")
		}
		fmt.Printf("%q\n", match)
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
			fmt.Printf("set: %+v\n", set)
			g.AddSet(set)
			set = Set{}
		}
	}

	return nil
}

type Set struct {
	Green, Red, Blue int // set holds counts for each 'cube'
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
		break
	}

	fmt.Printf("%v\n", games)
	fmt.Printf("found %d games\n", len(games))
	for _, g := range games {
		fmt.Printf("%+v\n", g)
	}
}
