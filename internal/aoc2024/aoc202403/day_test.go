package aoc202403

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sampleInput []byte

//go:embed sample2.txt
var sample2 []byte

//go:embed input.txt
var testInput []byte

func TestPart1(t *testing.T) {
	assert.Equal(t, int64(161), Part1(sampleInput))
	assert.Equal(t, int64(174960292), Part1(testInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, int64(48), Part2(sample2))
	assert.Equal(t, int64(56275602), Part2(testInput))
}
