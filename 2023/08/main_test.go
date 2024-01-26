package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput1 = []byte(`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`)

var exampleInput2 = []byte(`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`)

var exampleInput3 = []byte(`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`)

func readExample(t *testing.T, example []byte) (gm GhostMap) {
	t.Helper()

	err := gm.UnmarshalText(example)
	assert.NoError(t, nil, err)
	return
}

func TestGhostMapPart1(t *testing.T) {
	cases := []struct {
		Example       []byte
		ExpectedSteps int
	}{
		{exampleInput1, 2},
		{exampleInput2, 6},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("example%d", i+1), func(t *testing.T) {
			gm := readExample(t, tc.Example)
			assert.Equal(t, tc.ExpectedSteps, gm.TotalSteps())
		})
	}
}

func TestGhostMapPart2(t *testing.T) {
	gm := readExample(t, exampleInput3)
	assert.Equal(t, 6, gm.TotalStepsPart2())
}

func BenchmarkGhostMapPart1(b *testing.B) {
	var gm GhostMap
	err := gm.UnmarshalText(exampleInput1)
	if err != nil {
		b.Fail()
	}

	for i := 0; i < b.N; i++ {
		gm.TotalSteps()
	}
}

func BenchmarkGhostMapPart2(b *testing.B) {
	var gm GhostMap
	err := gm.UnmarshalText(exampleInput1)
	if err != nil {
		b.Fail()
	}

	for i := 0; i < b.N; i++ {
		gm.TotalStepsPart2()
	}
}
