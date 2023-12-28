package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"

	"github.com/kmullin/advent-of-code/2023/common"
)

var cardHeader = regexp.MustCompile(`^Card\s+(\d+): (.*)$`)

type cardReader struct {
	cards []card
}

type card struct {
	ID            int   // Id of the card
	winning, have []int // the actual card numbers
	matching      int   // the number of winning cards
}

func (cr *cardReader) UnmarshalText(text []byte) error {
	scanner := bufio.NewScanner(bytes.NewReader(text))
	for scanner.Scan() {
		match := cardHeader.FindSubmatch(scanner.Bytes())
		if len(match) != 3 {
			return fmt.Errorf("unable to parse line correctly %+v", match)
		}

		var card card
		card.ID, _ = strconv.Atoi(string(match[1]))

		// we have 2 sides to the parsing, split at most 2, use indexes directly
		for i, part := range bytes.SplitN(match[2], []byte("|"), 2) {
			nums := convertDigits(part)
			switch i {
			case 0: // winning
				card.winning = nums
			case 1: // have
				card.have = nums
			default:
				panic("should not happen / split too much")
			}
		}
		cr.cards = append(cr.cards, card)
		fmt.Println(card.String())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Value determines total value of all cards together
func (cr *cardReader) Value() (sum int) {
	for _, card := range cr.cards {
		sum += card.Value()
	}
	return
}

// RevisedValue calculates card point totals from Part 2 findings
func (cr *cardReader) RevisedValue() (sum int) {
	// cardDupes := make(map[int]int)
	for _, card := range cr.cards {
		fmt.Println(card.String())
	}
	return
}

func (c *card) MatchingCards() (count int) {
	for _, i := range c.have {
		if slices.Contains(c.winning, i) {
			count++
		}
	}
	return
}

// Value returns the point value of the card
func (c *card) Value() int {
	return int(math.Pow(2.0, float64(c.MatchingCards()-1)))
}

func (c *card) String() string {
	return fmt.Sprintf("Card %v: Winning %v / Have %v (%v) (%v matches)", c.ID, c.winning, c.have, c.Value(), c.MatchingCards())
}

func convertDigits(b []byte) (ints []int) {
	for _, num := range bytes.Split(b, []byte(" ")) {
		if len(num) == 0 {
			continue
		}
		// reasonably sure we're only dealing with numbers
		n, _ := strconv.Atoi(string(num))
		ints = append(ints, n)
	}
	return
}

func main() {
	var filename common.FileFlag
	flag.Var(&filename, "input-file", "what")
	flag.Parse()

	var cr cardReader
	err := cr.UnmarshalText(filename.Content)
	if err != nil {
		fmt.Printf("err unmarshaling text: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("total value of all cards: %v\n", cr.Value())
	fmt.Printf("total revised value of all cards: %v\n", cr.RevisedValue())
}
