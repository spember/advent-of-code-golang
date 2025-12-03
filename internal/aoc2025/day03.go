package aoc2025

import (
	"fmt"
	"strconv"

	"github.com/spember/advent-of-code-golang/pkg/aocutils/printer"
)

/*
Day 3 - Lobby
https://adventofcode.com/2025/day/3

*/

type Lobby struct {
}

func (l *Lobby) Part1(lines []string) int {
	return int(l.solve(lines, 2))
}

func (l *Lobby) Part2(lines []string) int64 {
	return int64(l.solve(lines, 12))
}

func (l *Lobby) solve(lines []string, digits int) int64 {
	var total int64 = 0
	for _, line := range lines {
		bank := l.convertTo(line)
		jolts := l.FindMaxNJoltage(bank, digits)
		printer.Ln("Found jolts ", jolts)
		total += jolts
	}
	return total
}

func (l *Lobby) FindMaxNJoltage(bank []int, digits int) int64 {
	if len(bank) < digits {
		fmt.Printf("Cannot find joltage for %v because not enough bank for digits %d\n", bank, digits)
		return 0
	}

	runes := ""
	printer.Ln("Bank is ", bank)
	// search space should be, for each digit, the leftmost space up until the -d

	var swallowedPos = 0

	for d := digits; d > 0; d-- {
		// find the largest number in bank len()-digits+d
		newP, val := l.findMaxInSlice(bank[swallowedPos : len(bank)-(d-1)])
		printer.F("found %d at position %d\n", val, newP)
		swallowedPos += newP + 1
		runes += strconv.Itoa(val)

	}

	int64Value, err := strconv.ParseInt(runes, 10, 64)
	if err != nil {
		panic(err)
	}
	return int64Value

}

func (l *Lobby) findMaxInSlice(bankSlice []int) (int, int) {
	var maxValue = 0
	var position = 0
	printer.Ln("Searching in ", bankSlice)
	for p, jolt := range bankSlice {
		if jolt > maxValue {
			maxValue = jolt
			position = p
		}
	}
	return position, maxValue
}

func (l *Lobby) convertTo(intLine string) []int {
	var result []int = make([]int, len(intLine))
	runes := []rune(intLine)
	for p, r := range runes {
		v, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		result[p] = v
	}
	return result
}
