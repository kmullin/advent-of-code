package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/kmullin/advent-of-code/2023/common"
)

var digitRe = regexp.MustCompile(`(\d+)`)

type Races []Race

type Race struct {
	Time     int
	Distance int
}

func (r *Races) UnmarshalText(text []byte) error {
	// we have 2 lines as input
	// Time
	// Distance
	// arbitrary lengh of line, but columns are separated by whitespace
	// and indicate more than one race
	matches := digitRe.FindAllSubmatch(text, -1)
	if len(matches) == 0 || len(matches)%2 != 0 {
		return errors.New("not an even number of digits in input")
	}
	line1, line2 := matches[:len(matches)/2], matches[len(matches)/2:]
	for i := 0; i < len(matches)/2; i++ {
		var race Race
		race.Time, _ = strconv.Atoi(string(line1[i][1]))
		race.Distance, _ = strconv.Atoi(string(line2[i][1]))
		*r = append(*r, race)
	}
	return nil
}

func (r *Races) NumberOfWinnableRaces() (possible int) {
	possible = 1
	for _, race := range *r {
		possible *= race.WinnableTimings()
	}
	return
}

func (r *Race) UnmarshalText(text []byte) (err error) {
	matches := digitRe.FindAllSubmatch(text, -1)
	if len(matches) == 0 || len(matches)%2 != 0 {
		return errors.New("not an even number of digits in input")
	}
	line1, line2 := matches[:len(matches)/2], matches[len(matches)/2:]

	var t, d []byte
	// append time digits together
	for i := 0; i < len(line1); i++ {
		t = append(t, line1[i][1]...)
	}
	r.Time, err = strconv.Atoi(string(t))
	if err != nil {
		return err
	}

	// now do distance
	for i := 0; i < len(line2); i++ {
		d = append(d, line2[i][1]...)
	}
	r.Distance, err = strconv.Atoi(string(d))
	if err != nil {
		return err
	}

	return nil
}

func (r *Race) WinnableTimings() (possibleWins int) {
	for msecs := 1; msecs <= r.Time; msecs++ {
		remainingRaceTime := r.Time - msecs
		// 1 millimeter per millisecond
		distanceTraveled := 1 * msecs * remainingRaceTime

		// if we traveled a greater distance than the current record
		if distanceTraveled > r.Distance {
			possibleWins++
		}
	}
	return
}

func main() {
	var filename common.FileFlag
	flag.Var(&filename, "input-file", "")
	flag.Parse()

	var races Races
	err := races.UnmarshalText(filename.Content)
	if err != nil {
		fmt.Printf("err unmarshaling text: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1: %+v\n", races.NumberOfWinnableRaces())

	var race Race
	err = race.UnmarshalText(filename.Content)
	if err != nil {
		fmt.Printf("err unmarshaling text: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2: %+v\n", race.WinnableTimings())
}
