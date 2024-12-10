package aoc202407

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

	assert.Equal(t, 3749, Part1(sampleInput))
	assert.Equal(t, 5030892084481, Part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, int64(11387), Part2(sampleInput))
	// failed: 24217859095482
	// this is not the right answer ... so no star for me. Some slight bug that I ran out of time looking for
	//assert.Equal(t, int64(24217859095482), Part2(input))
}
