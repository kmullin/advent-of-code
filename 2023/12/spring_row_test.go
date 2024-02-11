package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpringRow(t *testing.T) {
	t.Run("noerror", func(t *testing.T) {
		sr, err := newSpringRow([]byte(`???....###?.### 1,2,3`))
		assert.NoError(t, err)

		assert.Equal(t, 3, len(sr.damagedSprings))
		assert.Equal(t, 15, len(sr.data))
	})

	t.Run("error", func(t *testing.T) {
		_, err := newSpringRow([]byte(`??.#1,3,b`))
		assert.Error(t, err)
	})
	t.Run("error", func(t *testing.T) {
		_, err := newSpringRow([]byte(`??.# 1,3,b`))
		assert.Error(t, err)
	})
}

func TestSpringRowUnfold(t *testing.T) {
	sr, err := newSpringRow([]byte(`.# 1`))
	assert.NoError(t, err)

	sr = sr.Unfold(unfold)
	assert.Equal(t, ".#?.#?.#?.#?.#", sr.data)
	assert.Len(t, sr.damagedSprings, 5)
}
