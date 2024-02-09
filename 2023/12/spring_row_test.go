package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpringRow(t *testing.T) {
	sr, err := newSpringRow([]byte(`???....###?.### 1,2,3`))
	assert.NoError(t, err)

	assert.Equal(t, 3, len(sr.damagedSprings), "should have 3 groupings of damaged springs")
	assert.Equal(t, 6, len(sr.groups), "should have 6 groupings of contiguous springs")
}
