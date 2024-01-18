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
	SeedsRe  = regexp.MustCompile(`^seeds:.*$`)
	MapStart = regexp.MustCompile(`^([\w-]+) map:$`)
)

// mapOrder is the order of the maps presented in the input
var mapOrder = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

type Almanac struct {
	Seeds []int
	Maps  map[string][]NumRange
}

type NumRange struct {
	DestinationStart int
	SourceStart      int
	Length           int
}

func (a *Almanac) UnmarshalText(text []byte) error {
	a.Maps = make(map[string][]NumRange)

	var lineNum int
	var mapName string

	scanner := bufio.NewScanner(bytes.NewReader(text))
	for scanner.Scan() {
		lineNum++
		if len(scanner.Bytes()) == 0 {
			// we have a blank line, reset map name
			mapName = ""
			continue
		}

		// seeds header, numbers are in a single line separated by spaces
		if lineNum == 1 && SeedsRe.Match(scanner.Bytes()) {
			s := bytes.Split(scanner.Bytes(), []byte(" "))
			for _, b := range s[1:] {
				i, _ := strconv.Atoi(string(b))
				a.Seeds = append(a.Seeds, i)
			}
			continue
		}

		if mapName == "" {
			if matches := MapStart.FindSubmatch(scanner.Bytes()); len(matches) == 2 {
				mapName = string(matches[1])
				continue
			}
		}

		// we are between map lines, need to read digits
		s := bytes.SplitN(scanner.Bytes(), []byte(" "), 3)
		if len(s) != 3 {
			return errors.New("weird range input length != 3")
		}
		if mapName == "" {
			// we dun goofed
			return errors.New("mapname is empty")
		}

		var n NumRange
		n.DestinationStart, _ = strconv.Atoi(string(s[0]))
		n.SourceStart, _ = strconv.Atoi(string(s[1]))
		n.Length, _ = strconv.Atoi(string(s[2]))

		a.Maps[mapName] = append(a.Maps[mapName], n)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (a *Almanac) Lookup(mapName string, n int) int {
	if nums, ok := a.Maps[mapName]; ok {
		for _, num := range nums {
			// if given number is in range of source
			// then we have a number that has a mapping
			// calculate offset and add offset to destinationStart
			if n >= num.SourceStart && n < num.SourceStart+num.Length {
				offset := n - num.SourceStart
				return num.DestinationStart + offset
			}
		}
	}

	return n
}

func (a *Almanac) LowestLocation() (lowest int) {
	for _, seed := range a.Seeds {
		num := seed
		for _, mapName := range mapOrder {
			num = a.Lookup(mapName, num)
		}
		if num < lowest || lowest == 0 {
			lowest = num
		}
	}
	return
}

func (a *Almanac) LowestLocationP2() (lowest int) {
	// treat seeds as a range of numbers
	for i := 0; i < len(a.Seeds); i += 2 {
		for seed := a.Seeds[i]; seed < a.Seeds[i]+a.Seeds[i+1]; seed++ {
			num := seed
			for _, mapName := range mapOrder {
				num = a.Lookup(mapName, num)
			}
			if num < lowest || lowest == 0 {
				lowest = num
			}
		}
	}
	return
}

func main() {
	var filename common.FileFlag
	flag.Var(&filename, "input-file", "")
	flag.Parse()

	var a Almanac
	err := a.UnmarshalText(filename.Content)
	if err != nil {
		fmt.Printf("err unmarshaling text: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Lowest location number from seeds (Part1): %d\n", a.LowestLocation())
	fmt.Printf("Lowest location number from seeds (Part2): %d\n", a.LowestLocationP2())
}
