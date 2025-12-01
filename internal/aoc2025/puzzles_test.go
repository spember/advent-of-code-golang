package aoc2025

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/spember/advent-of-code-golang/pkg/aocutils/parseto"
	"github.com/spember/advent-of-code-golang/pkg/aocutils/printer"
)

//go:embed testdata/day01_sample_1.txt
var day01Sample []byte

//go:embed testdata/day01_input_1.txt
var day01Input1 []byte

func TestDay01(t *testing.T) {

	solver := &SecretEntrance{}

	printer.Enable()

	t.Run("Part 1 - Sample Input", func(t *testing.T) {
		assert.Equal(t, 3, solver.Part1(day01Sample))
	})

	t.Run("Part 1 - Solver Input", func(t *testing.T) {
		assert.Equal(t, 1034, solver.Part1(day01Input1))
	})
	
	t.Run("Part 2 Side Cases", func(t *testing.T) {
		assert.Equal(t, 1, solver.Part2([]string{
			"L50",
			"R50",
		}))
		// how is this screwing up?
		// start at 50, go up by 50. count that we landed at zero (100)
		// go back down by 50, to 50
		assert.Equal(t, 1, solver.Part2([]string{
			"R50",
			"L50",
		}))

		assert.Equal(t, 1, solver.Part2([]string{
			"L50",
			"L50",
		}))

		assert.Equal(t, 1, solver.Part2([]string{
			"R50",
			"R50",
		}))

		assert.Equal(t, 2, solver.Part2([]string{
			"L150",
			"L50",
		}))

		assert.Equal(t, 2, solver.Part2([]string{
			"R150",
			"L50",
		}))

		assert.Equal(t, 2, solver.Part2([]string{
			"L150",
			"R50",
		}))

		assert.Equal(t, 2, solver.Part2([]string{
			"R150",
			"R50",
		}))

	})

	t.Run("Part 2 - Solver Input", func(t *testing.T) {
		printer.Disable()
		defer printer.Enable()
		assert.Equal(t, 6, solver.Part2(parseto.Lines(day01Sample)))
		assert.Equal(t, 6173, solver.Part2(parseto.Lines(day01Input1)))
		//6173 is too high, so is 6277. Got this wrong

	})
}
