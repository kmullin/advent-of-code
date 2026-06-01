package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/kmullin/advent-of-code/internal/cli"
	"github.com/rs/zerolog/log"
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

func (g *game) MinimumPossible() (cs cubeSet) {
	for _, s := range g.Sets {
		if s.Red > cs.Red {
			cs.Red = s.Red
		}
		if s.Green > cs.Green {
			cs.Green = s.Green
		}
		if s.Blue > cs.Blue {
			cs.Blue = s.Blue
		}
	}
	return
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

func part1(b []byte) (any, error) {
	games, err := getGames(b)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("sum of games that were possible: %d", gamesPossible(games)), nil
}

func part2(b []byte) (any, error) {
	games, err := getGames(b)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("sum of the power of sets: %d", powerOfMinimumSets(games)), nil
}

func getGames(b []byte) ([]game, error) {
	var games []game
	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		var g game
		err := g.UnmarshalText(scanner.Bytes())
		if err != nil {
			return nil, err
		}
		games = append(games, g)
	}

	log.Debug().Msgf("found %d games", len(games))
	return games, nil
}

func main() {
	if err := cli.NewCmd(2023, 2, part1, part2).Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func gamesPossible(games []game) (sum int) {
	for _, g := range games {
		if g.IsPossible(12, 13, 14) {
			sum += g.ID
		}
	}
	return
}

func powerOfMinimumSets(games []game) (power int) {
	for _, g := range games {
		cs := g.MinimumPossible()
		power += cs.Red * cs.Green * cs.Blue
	}
	return
}
