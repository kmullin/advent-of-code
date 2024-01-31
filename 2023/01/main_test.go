package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalibrationValue(t *testing.T) {
	cases := []struct {
		input  string
		number int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			n := getDigit(tc.input)
			assert.Equal(t, tc.number, n)
		})
	}
}

func TestCalibrationValue2(t *testing.T) {
	cases := []struct {
		input  string
		number int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			n := getDigitWords(tc.input)
			assert.Equal(t, tc.number, n)
		})
	}
}
