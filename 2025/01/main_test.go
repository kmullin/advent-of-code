package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamplePart1(t *testing.T) {
	cases := []struct {
		rotation Rotation
		expected int
	}{
		{Rotation{Left, 68}, 82},
		{Rotation{Left, 30}, 52},
		{Rotation{Right, 48}, 0},
		{Rotation{Left, 5}, 95},
		{Rotation{Right, 60}, 55},
		{Rotation{Left, 55}, 0},
		{Rotation{Left, 1}, 99},
		{Rotation{Left, 99}, 0},
		{Rotation{Right, 14}, 14},
		{Rotation{Left, 82}, 32},
	}

	d := NewDial(0, 99, 50)

	for _, tc := range cases {
		assert.Equal(t, tc.expected, d.Rotate(tc.rotation))
		t.Logf("rotated %v, new pos %v", tc.rotation, d.pos)
	}
}

func TestExamplePart2(t *testing.T) {
	cases := []struct {
		rotation Rotation
		expected int
	}{
		{Rotation{Left, 68}, 1},
		{Rotation{Left, 30}, 0},
		{Rotation{Right, 48}, 1},
		{Rotation{Left, 5}, 0},
		{Rotation{Right, 60}, 1},
		{Rotation{Left, 55}, 1},
		{Rotation{Left, 1}, 0},
		{Rotation{Left, 99}, 1},
		{Rotation{Right, 14}, 0},
		{Rotation{Left, 82}, 1},
	}

	d := NewDial(0, 99, 50)

	num_zeros := 0
	for _, tc := range cases {
		//assert.Equal(t, tc.expected, d.RotateCount(tc.rotation))
		n := d.RotateCount(tc.rotation)
		num_zeros += n
		t.Logf("rotated %v, num zeros %v", tc.rotation, n)
		assert.Equal(t, tc.expected, n)
	}

	const total_expected = 6
	assert.Equal(t, total_expected, num_zeros)
}
