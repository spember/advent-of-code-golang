package aoc202402

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 2, Part1(sampleInput))
	assert.Equal(t, 242, Part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 4, Part2(sampleInput))
	assert.Equal(t, 311, Part2(input))
}
