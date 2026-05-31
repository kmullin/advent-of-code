package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	t.Skip("need to mock this test")
	c := NewClient("foobar", t.TempDir())

	b, err := c.GetInput(2025, 7)
	assert.NoError(t, err)

	t.Logf("%s", b)
}
