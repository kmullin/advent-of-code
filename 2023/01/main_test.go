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
		n, err := GetDigit(tc.input)
		assert.Nil(t, err)
		assert.Equal(t, tc.number, n)
	}
}
