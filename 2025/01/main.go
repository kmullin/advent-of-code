package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/kmullin/advent-of-code/internal/util"
)

const (
	Left = iota
	Right
)

type Rotation struct {
	Direction int
	Count     int
}

func (r Rotation) String() string {
	var dir string
	switch r.Direction {
	case Left:
		dir = "L"
	case Right:
		dir = "R"
	}
	return fmt.Sprintf("%s%d", dir, r.Count)
}

func ReadInput(r io.Reader) ([]Rotation, error) {
	var rotations []Rotation

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		b := scanner.Bytes()

		var r Rotation
		switch v := b[0]; v {
		case 'L':
			r.Direction = Left
		case 'R':
			r.Direction = Right
		default:
			log.Fatalf("unknown first byte %q", v)
		}

		n, err := strconv.Atoi(string(b[1:]))
		if err != nil {
			return nil, fmt.Errorf("unable to convert: %w", err)
		}

		r.Count = n
		rotations = append(rotations, r)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rotations, nil
}

// Dial is the vault dial
type Dial struct {
	steps int // how many clicks there are on the dial
	pos   int // the current position
}

// NewDial returns a new dial starting at pos
func NewDial(start, end, pos int) *Dial {
	return &Dial{
		steps: end - start + 1,
		pos:   pos,
	}
}

// Rotate rotates the dial with a given rotation, returns the resulting position
func (d *Dial) Rotate(r Rotation) int {
	count := r.Count
	// turning left is the same number of right rotations as steps-n
	if r.Direction == Left {
		count = d.steps - r.Count
	}

	d.pos = (d.pos + count) % d.steps
	return d.pos
}

// RotateCount rotates the dial with a given rotation, and keeps track of how many times it clicks on 0
func (d *Dial) RotateCount(r Rotation) int {
	num_zeros := 0
	for range r.Count {
		if r.Direction == Left {
			d.pos = (d.pos - 1) % d.steps
		} else {
			d.pos = (d.pos + 1) % d.steps
		}

		if d.pos == 0 {
			num_zeros += 1
		}
	}

	return num_zeros
}

func main() {
	rotations, err := ReadInput(util.InputReader())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v total rotations\n", len(rotations))

	d := NewDial(0, 99, 50)

	// part 1
	zero_pos := 0
	for _, r := range rotations {
		pos := d.Rotate(r)
		if pos == 0 {
			zero_pos += 1
		}
	}
	fmt.Printf("number of zeros: %v\n", zero_pos)

	// reset
	d = NewDial(0, 99, 50)

	// part 2
	num_zeros := 0
	for _, r := range rotations {
		num_zeros += d.RotateCount(r)
	}
	fmt.Printf("number of times hit zero: %v\n", num_zeros)
}
