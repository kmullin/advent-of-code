package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
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

func GetDigitWords(input string) (int, error) {
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

	return strconv.Atoi(fmt.Sprintf("%d%d", digits[0], digits[1]))
}

func GetDigit(input string) (int, error) {
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
	return strconv.Atoi(string(digits))
}

func main() {
	var filename common.FileFlag
	flag.Var(&filename, "input-file", "what")
	flag.Parse()

	answer, err := Scan(bytes.NewReader(filename.Content), GetDigitWords)
	if err != nil {
		fmt.Printf("err encountered: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("the answer is: %v\n", answer)
}

func Scan(r io.Reader, f func(string) (int, error)) (int, error) {
	var answer int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n, err := f(scanner.Text())
		if err != nil {
			return -1, err
		}
		answer += n
	}

	return answer, nil
}
