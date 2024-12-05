package aoc202405

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/sample.txt
var sampleInput []byte

//go:embed testdata/input.txt
var input []byte

func TestPart1(t *testing.T) {
	assert.Equal(t, 143, Part1(sampleInput))
	assert.Equal(t, 5651, Part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 123, Part2(sampleInput))
	assert.Equal(t, 4743, Part2(input))
}
