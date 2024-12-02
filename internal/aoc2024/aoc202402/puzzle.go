package aoc202402

import (
	"cmp"
	_ "embed"
	"math"
	"strings"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

/*
--- Day 02: Red Nosed Reports

I used a struct to capture the Report concept, but I needn't have bothered.
*/

//go:embed sample.txt
var sampleInput []byte

//go:embed input.txt
var input []byte

func Part1(input []byte) int {
	lines := aocutils.ParseLines(input)
	var numSafe = 0
	for _, line := range lines {
		report := NewReport(line)
		if report.IsSafe() {
			numSafe++
		}
	}

	return numSafe
}

func Part2(input []byte) int {
	lines := aocutils.ParseLines(input)
	var numSafe = 0
	// for each line, try removing one element and see if the report is safe
	for _, line := range lines {
		baseReport := NewReport(line)
		if baseReport.IsSafe() {
			numSafe++
			continue
		}
		levels := baseReport.GetLevels()
		for i := 0; i < len(levels); i++ {
			var modifiedLevels = append([]int{}, levels...)
			modifiedLevels = append(modifiedLevels[:i], modifiedLevels[i+1:]...)
			modifiedReport := &RawReport{Levels: modifiedLevels}
			if modifiedReport.IsSafe() {
				numSafe++
				break
			}
		}
	}

	return numSafe
}

func NewReport(line string) Report {
	chunks := strings.Split(line, " ")
	return &RawReport{Levels: aocutils.Map(chunks, aocutils.IntOrPanic)}
}

type Report interface {
	IsSafe() bool
	GetLevels() []int
}

type RawReport struct {
	Levels []int
}

func (r *RawReport) GetLevels() []int {
	return r.Levels
}

func (r *RawReport) IsSafe() bool {
	_, isBad := r.detectBadLevel()
	return !isBad
}

func (r *RawReport) detectBadLevel() (int, bool) {
	var isIncreasing = cmp.Less(r.Levels[0], r.Levels[1])
	for i := 0; i < len(r.Levels)-1; i++ {
		// first check the absolute value of the difference between the two levels

		diff := math.Abs(float64(r.Levels[i+1] - r.Levels[i]))
		if diff < 1 || diff > 3 {
			//fmt.Println("Not safe due to difference")
			return i, true
		}
		// then check if the levels are increasing or decreasing. how to check if a sequence is always increasing or always decreasing?
		if isIncreasing != cmp.Less(r.Levels[i], r.Levels[i+1]) {
			//fmt.Println("Not safe due to direction")
			return i, true
		}

	}
	return -1, false
}
