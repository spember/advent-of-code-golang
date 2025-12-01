package aoc2025

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/day01_sample_1.txt
var day01Sample []byte

//go:embed testdata/day01_input_1.txt
var day01Input1 []byte

func TestDay01(t *testing.T) {

	solver := &SecretEntrance{}

	t.Run("Part 1 - Sample Input", func(t *testing.T) {
		assert.Equal(t, 3, solver.Part1(day01Sample))
	})

	t.Run("Part 1 - Solver Input", func(t *testing.T) {
		assert.Equal(t, 1034, solver.Part1(day01Input1))
	})

	t.Run("Part 2", func(t *testing.T) {
		assert.Equal(t, 6, solver.Part2(day01Sample))
		assert.Equal(t, 6173, solver.Part2(day01Input1))
		//6173 is too high, so is 6277. Got this wrong

	})
}
