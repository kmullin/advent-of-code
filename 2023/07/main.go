package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/kmullin/advent-of-code/2023/common"
)

type Hands []Hand

type Hand struct {
	Cards string
	Bid   int
}

type cardType int

const handSize = 5

const (
	// Card Types in reverse order from lowest rank to highest
	HighCard cardType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

// ValidCards is a mapping of cards to relative strength
var ValidCards = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	// 'J' becomes lowest value in Part 2
}

func (h *Hands) UnmarshalText(text []byte) error {
	scanner := bufio.NewScanner(bytes.NewReader(text))
	for scanner.Scan() {
		sides := strings.Split(scanner.Text(), " ")
		if len(sides) != 2 {
			return errors.New("invalid line, expected one split")
		}
		if len(sides[0]) != handSize {
			return errors.New("invalid hand, did not find 5 cards")
		}

		var hand Hand
		hand.Cards = sides[0]
		hand.Bid, _ = strconv.Atoi(sides[1])
		*h = append(*h, hand)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// TotalWinnings determines the total winning score of all hands.
// multiply each hand's bid with its rank
func (h Hands) TotalWinnings(part int) (total int) {
	switch part {
	case 1: // part 1
		// XXX: this changes global state
		ValidCards['J'] = 10
	case 2: // part 2
		// 'J' is now the least valuable card, and counts as a wildcard
		// assert it lowest on the priority list
		// XXX: this changes global state
		ValidCards['J'] = 0
	default:
		panic(fmt.Sprintf("parts unknown: %d", part)) // RIP A.B.
	}

	sort.Slice(h, h.Less)

	for i, hand := range h {
		// we've sorted in descending order, so rank is calculated
		rank := len(h) - i
		total += hand.Bid * rank
	}
	return
}

// Less reports whether the element with index i must sort before the element with index j.
func (h Hands) Less(i, j int) bool {
	// if we have the same type, do second ordering
	if h[i].Type() == h[j].Type() {
		for ii := 0; ii < handSize; ii++ {
			if h[i].Cards[ii] == h[j].Cards[ii] {
				// we have the same card in the same position
				continue
			}
			return ValidCards[rune(h[i].Cards[ii])] > ValidCards[rune(h[j].Cards[ii])]
		}
	}
	return h[i].Type() > h[j].Type()
}

// Five of a kind, where all five cards have the same label: AAAAA
// Four of a kind, where four cards have the same label and one card has a different label: AA8AA
// Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
// Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// High card, where all cards' labels are distinct: 23456
func (h *Hand) Type() cardType {
	m := make(map[rune]int)
	for _, r := range h.Cards {
		m[r]++
	}

	// XXX: uses global state
	if ValidCards['J'] == 0 {
		if jokers := m['J']; jokers > 0 {
			// add count to highest card, delete J
			m[getMostCard(m)] += jokers
			delete(m, 'J')
		}
	}

	switch len(m) {
	case 1: // all cards are the same
		return FiveOfAKind
	case 2:
		if isOfAKind(m, 4) {
			return FourOfAKind
		}
		return FullHouse
	case 3:
		if isOfAKind(m, 3) {
			return ThreeOfAKind
		}
		return TwoPair
	case 4:
		return OnePair
	case 5: // all cards are unique
		return HighCard
	default:
		return -1
	}
}

func getMostCard(m map[rune]int) rune {
	var most struct {
		card  rune
		count int
	}

	for r, c := range m {
		if c >= most.count && r != 'J' {
			if c == most.count && !compareRank(r, most.card) {
				continue
			}
			most.card = r
			most.count = c
		}
	}

	return most.card
}

// compareRank compares the value of each card, returns true if value of a > b
func compareRank(a, b rune) bool {
	return ValidCards[a] > ValidCards[b]
}

// isOfAKind checks to see if we have n of the same type in s
func isOfAKind(m map[rune]int, n int) bool {
	for _, c := range m {
		if c == n {
			return true
		}
	}
	return false
}

func main() {
	flag.Parse()

	var h Hands
	err := common.FileFlag(&h)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d hands\n", len(h))
	for i := 1; i <= 2; i++ {
		fmt.Printf("Total Winnings Part %d: %d\n", i, h.TotalWinnings(i))
	}
}
