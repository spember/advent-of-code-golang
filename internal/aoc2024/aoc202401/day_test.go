package aoc202401

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 11, Part1(sampleInput))
	assert.Equal(t, 1834060, Part1(myInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 31, Part2(sampleInput))
	assert.Equal(t, 21607792, Part2(myInput))
}
