package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/kmullin/advent-of-code/2023/common"
)

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

	var answer int

	scanner := bufio.NewScanner(bytes.NewReader(filename.Content))
	for scanner.Scan() {
		fmt.Println(scanner.Text())

		n, err := GetDigit(scanner.Text())
		if err != nil {
			fmt.Printf("err encountered: %v\n", err)
			os.Exit(1)
		}
		answer += n
	}

	fmt.Printf("the answer is: %v\n", answer)
}
