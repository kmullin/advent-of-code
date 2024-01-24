package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func readInput(t *testing.T) Hands {
	t.Helper()

	var h Hands
	err := h.UnmarshalText([]byte(exampleInput))
	assert.Nil(t, err)
	assert.Len(t, h, handSize)
	return h
}

func TestScoring(t *testing.T) {
	h := readInput(t)

	cases := []struct {
		Part     int
		Expected int
	}{
		{1, 6440},
		{2, 5905},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Part%d", tc.Part), func(t *testing.T) {
			assert.Equal(t, tc.Expected, h.TotalWinnings(tc.Part))
		})
	}
}

func TestHands(t *testing.T) {
	cases := []struct {
		Name     string
		Hand     string
		Expected cardType
	}{
		{"FiveOfAKind", "AAAAA", FiveOfAKind},
		{"FourOfAKind", "AA8AA", FourOfAKind},
		{"FullHouse", "23332", FullHouse},
		{"ThreeOfAKind", "TTT98", ThreeOfAKind},
		{"TwoPair", "23432", TwoPair},
		{"OnePair", "A23A4", OnePair},
		{"HighCard", "23456", HighCard},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			h := Hand{tc.Hand, -1}
			assert.Equal(t, tc.Expected, h.Type())
		})
	}
}

func TestExample(t *testing.T) {
	h := readInput(t)

	t.Run("Part1Rules", func(t *testing.T) {
		values := []cardType{
			OnePair,
			ThreeOfAKind,
			TwoPair,
			TwoPair,
			ThreeOfAKind,
		}
		ValidCards['J'] = 10 // XXX: global state bad
		for i := 0; i < len(h); i++ {
			assert.Equal(t, values[i], h[i].Type())
		}
	})

	t.Run("Part2Rules", func(t *testing.T) {
		values := []cardType{
			OnePair,
			FourOfAKind,
			TwoPair,
			FourOfAKind,
			FourOfAKind,
		}
		ValidCards['J'] = 0 // XXX: global state bad
		for i := 0; i < len(h); i++ {
			assert.Equal(t, values[i], h[i].Type())
		}
	})
}
