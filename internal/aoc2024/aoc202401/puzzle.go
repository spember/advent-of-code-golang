package aoc202401

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

//go:embed sample.txt
var sampleInput []byte

//go:embed input.txt
var myInput []byte

func Part1(input []byte) int {

	left, right := buildListsFromInput(input)

	slices.Sort(left)
	slices.Sort(right)

	var total = 0
	var pos = 0
	fmt.Println(left)
	fmt.Println(right)

	for {
		if pos >= len(left) {
			break
		}

		distance := right[pos] - left[pos]
		total += int(math.Abs(float64(distance)))

		pos++
	}

	return total
}

func Part2(input []byte) int {
	left, right := buildListsFromInput(input)

	var total = 0

	// for each number in the left list, count how many times it appears in the right list and multiply the left value
	//by right occurances
	for _, leftValue := range left {
		total += leftValue * aocutils.SliceCount(right, leftValue)
	}
	return total
}

func buildListsFromInput(input []byte) ([]int, []int) {
	lines := aocutils.ParseLines(input)
	var left = make([]int, len(lines))
	var right = make([]int, len(lines))

	for i, line := range lines {
		chunks := strings.Split(line, " ")
		left[i] = aocutils.IntOrPanic(chunks[0])
		right[i] = aocutils.IntOrPanic(chunks[len(chunks)-1])
	}

	return left, right
}
