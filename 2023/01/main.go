package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/kmullin/advent-of-code/2023/common"
)

var intMap = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

type Trebuchet struct {
	input string
}

func (t *Trebuchet) UnmarshalText(text []byte) error {
	t.input = string(text)
	return nil
}

func (t Trebuchet) forEachLine(f func(string) int) (answer int) {
	scanner := bufio.NewScanner(strings.NewReader(t.input))
	for scanner.Scan() {
		answer += f(scanner.Text())
	}
	return
}

func (t Trebuchet) GetDigitWords() int {
	return t.forEachLine(getDigitWords)
}

func (t Trebuchet) GetDigit() int {
	return t.forEachLine(getDigit)
}

func getDigitWords(input string) int {
	first := -1
	last := -1
	digits := make([]int, 2)
	for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		// try both words and digits
		for _, w := range []string{strconv.Itoa(i), intMap[i]} {
			n := strings.Index(input, w)
			if n != -1 && (n < first || first == -1) {
				first = n
				digits[0] = i
			}
			n = strings.LastIndex(input, w)
			if n != -1 && n > last {
				last = n
				digits[1] = i
			}
		}
	}

	n, _ := strconv.Atoi(fmt.Sprintf("%d%d", digits[0], digits[1]))
	return n
}

func getDigit(input string) int {
	var digits []rune
	var lastDigit rune
	for i, a := range input {
		if unicode.IsDigit(a) {
			if len(digits) == 0 {
				digits = append(digits, a)
			} else {
				lastDigit = a
			}
		}

		if i == len(input)-1 {
			digits = append(digits, lastDigit)
		}
	}
	// check for possible only digit, first digit is also the last digit
	if len(digits) == 2 && digits[1] == 0 {
		digits[1] = digits[0]
	}

	n, _ := strconv.Atoi(string(digits))
	return n
}

func main() {
	flag.Parse()

	var t Trebuchet
	err := common.FileFlag(&t)
	if err != nil {
		fmt.Printf("err encountered: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("the answer is: %v\n", t.GetDigit())
	fmt.Printf("the answer is (Part 2): %v\n", t.GetDigitWords())
}
