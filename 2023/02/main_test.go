package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleInput = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestGameParsing(t *testing.T) {
	var games []game

	scanner := bufio.NewScanner(strings.NewReader(exampleInput))
	for scanner.Scan() {
		var g game
		err := g.UnmarshalText(scanner.Bytes())
		assert.Nil(t, err)
		games = append(games, g)
	}

	for _, g := range games {
		t.Logf("game %d possible: %v", g.ID, g.IsPossible(12, 13, 14))
	}
}
