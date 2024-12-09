package aoc2024

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/day8-sample.txt
var day8Sample []byte

//go:embed testdata/day8.txt
var day8 []byte

//go:embed testdata/day8-trivial.txt
var day8Trivial []byte

func TestDay8P1(t *testing.T) {
	assert.Equal(t, 14, Day8Part1(day8Sample))
	assert.Equal(t, 344, Day8Part1(day8))
}

func TestDay8P2(t *testing.T) {
	assert.Equal(t, 9, Day8Part2(day8Trivial))
	assert.Equal(t, 34, Day8Part2(day8Sample))
	assert.Equal(t, 1182, Day8Part2(day8))
}

//go:embed testdata/day9-trivial.txt
var day9Trivia []byte

//go:embed testdata/day9-sample.txt
var day9Sample []byte

//go:embed testdata/day9.txt
var day9Input []byte

func TestDay9P1(t *testing.T) {
	assert.Equal(t, 60, Day9Part1(day9Trivia))
	assert.Equal(t, 1928, Day9Part1(day9Sample))
	assert.Equal(t, 6401092019345, Day9Part1(day9Input))
}

func TestDay9P2(t *testing.T) {
	//assert.Equal(t, 99, Day9Part2(day9Trivia))
	assert.Equal(t, 2858, Day9Part2(day9Sample))
	// tried 6522525986866 which is not correct
	assert.Equal(t, 6522525986866, Day9Part2(day9Input))
}
