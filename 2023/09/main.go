package main

import (
	"bufio"
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/kmullin/advent-of-code/internal/cli"
	"github.com/rs/zerolog/log"
)

// OASIS Oasis And Sand Instability Sensor
type OASIS struct {
	History [][]int
}

func (o *OASIS) UnmarshalText(text []byte) error {
	scanner := bufio.NewScanner(bytes.NewReader(text))
	for scanner.Scan() {
		var history []int
		for _, s := range strings.Split(scanner.Text(), " ") {
			n, _ := strconv.Atoi(s)
			history = append(history, n)
		}
		o.History = append(o.History, history)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Extrapolate extrapolates the history from the current History, takes 1 or 2 as inputs for part1/2
func (o *OASIS) Extrapolate(part int) (answer int) {
	for _, nums := range o.History {
		switch part {
		case 1:
		case 2:
			slices.Reverse(nums)
		default:
			panic(fmt.Sprintf("parts unknown %v", part))
		}
		var stack []int
		for !allZeros(nums) {
			// record the last number of this iteration
			stack = append(stack, nums[len(nums)-1])
			nums = findSteps(nums)
		}

		answer += stackTotal(stack)
	}
	return
}

// stackTotal adds up all values in the slice
func stackTotal(nums []int) (sum int) {
	for _, n := range nums {
		sum += n
	}
	return
}

// allZeros determines if our slice is all zeros
func allZeros(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

// findSteps finds the amount of 'steps' between each number in nums
func findSteps(nums []int) (steps []int) {
	for i := 1; i < len(nums); i++ {
		steps = append(steps, nums[i]-nums[i-1])
	}
	return
}

func part1(b []byte) (any, error) {
	var o OASIS
	if err := o.UnmarshalText(b); err != nil {
		return nil, err
	}

	return fmt.Sprintf("History (%v) Extrapolated: %v",
		len(o.History),
		o.Extrapolate(1)), nil
}

func part2(b []byte) (any, error) {
	var o OASIS
	if err := o.UnmarshalText(b); err != nil {
		return nil, err
	}

	return fmt.Sprintf("History (%v) Extrapolated: %v",
		len(o.History),
		o.Extrapolate(2)), nil
}

func main() {
	if err := cli.NewCmd(2023, 9, part1, part2).Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
