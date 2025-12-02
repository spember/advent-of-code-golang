package aoc2025

import (
	_ "embed"
	"fmt"
	"strings"
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
		assert.Equal(t, 6166, solver.Part2(parseto.Lines(day01Input1)))
		
	})
}

//go:embed testdata/day02_sample_1.txt
var day02Sample []byte

//go:embed testdata/day02_input_1.txt
var day02Input []byte

func TestDay02(t *testing.T) {

	solver := &GiftShop{}

	strings.Split(parseto.Lines(day02Sample)[0], ",")

	t.Run("Part 1 - Sample Input", func(t *testing.T) {
		printer.Enable()
		assert.Equal(t, 1227775554, solver.Solve(solver.ParseLine(parseto.Lines(day02Sample)[0]), true))
		printer.Disable()
		defer printer.Enable()
		assert.Equal(t, 40398804950, solver.Solve(solver.ParseLine(parseto.Lines(day02Input)[0]), true))
	})

	type complexCase struct {
		Input  int
		Result bool
	}

	complexSolverInput := []complexCase{
		{Input: 11, Result: false},
		{Input: 123, Result: true},
		{Input: 22, Result: false},
		{Input: 111, Result: false},
		{Input: 1188511885, Result: false},
		{Input: 12341234, Result: false},
		{Input: 12341235, Result: true},
	}
	t.Run("Random test", func(t *testing.T) {
		for _, c := range complexSolverInput {
			t.Run(fmt.Sprintf("Testing %d should be %v", c.Input, c.Result), func(t *testing.T) {
				assert.Equal(t, c.Result, solver.RepeatingIsIdValid(c.Input))
			})
		}
	})

	t.Run("Part 2 - Solver Input", func(t *testing.T) {
		assert.Equal(t, 4174379265, solver.Solve(solver.ParseLine(parseto.Lines(day02Sample)[0]), false))
		printer.Disable()
		defer printer.Enable()
		assert.Equal(t, 65794984339, solver.Solve(solver.ParseLine(parseto.Lines(day02Input)[0]), false))
	})

}
