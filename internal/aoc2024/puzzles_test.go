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
