package aoc202406

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
	assert.Equal(t, 41, Part1(sampleInput))
	assert.Equal(t, 5212, Part1(input))
}
