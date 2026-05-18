package common

import (
	"log"
	"strconv"
)

func MustAtoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatalf("failed to convert %q: %v", a, err)
	}

	return i
}
