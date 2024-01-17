package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleInput = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func TestCardReader(t *testing.T) {
	var cr cardReader
	err := cr.UnmarshalText([]byte(exampleInput))
	assert.Nil(t, err)

	tc := []struct {
		cardID int
		points int
	}{
		{1, 8}, // card 1 is worth 8 points
		{2, 2},
		{3, 2},
		{4, 1},
		{5, 0},
		{6, 0},
	}

	for _, c := range tc {
		t.Run(fmt.Sprintf("Card %d", c.cardID), func(t *testing.T) {
			// slice stars with 0, cards start with 1
			assert.Equal(t, c.points, cr.cards[c.cardID-1].Value(), "card has incorrect value")
		})
	}
}

func TestCardReaderRevised(t *testing.T) {
	var cr cardReader
	err := cr.UnmarshalText([]byte(exampleInput))
	assert.Nil(t, err)

	// slice stars with 0, cards start with 1
	assert.Equal(t, 30, cr.RevisedValue(), "total scorecards has incorrect value")
}
