package aoc202404

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/trivial.txt
var trivialInput []byte

//go:embed testdata/sample.txt
var sampleInput []byte

//go:embed testdata/xmassample.txt
var xmasSample []byte

//go:embed testdata/input.txt
var monster []byte

func TestPart1(t *testing.T) {
	assert.Equal(t, 4, Part1(trivialInput))
	assert.Equal(t, 18, Part1(sampleInput))
	assert.Equal(t, 2613, Part1(monster))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 9, Part2(xmasSample))
	assert.Equal(t, 1905, Part2(monster))
}
