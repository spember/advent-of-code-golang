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
	var total = 0
	for _, line := range lines {
		bank := l.convertTo(line)
		jolts := l.P1Joltage(bank)
		printer.Ln("Found jolts ", jolts)
		total += jolts

	}
	return total
}

func (l *Lobby) P1Joltage(joltage []int) int {
	// find the largest number from 0 to len-2 as n1
	// then find the largest number from position(n1) to len-1)

	var n1 int = joltage[0]
	var n1pos = 0

	for i := 0; i < len(joltage)-1; i++ {
		if joltage[i] > n1 {
			n1 = joltage[i]
			n1pos = i
		}
	}
	//fmt.Printf("n1 joltage is %d at position %d\n", n1, n1pos)

	var n2 int = 0
	//var n2pos = 0
	for i := n1pos + 1; i < len(joltage); i++ {
		if joltage[i] > n2 {
			n2 = joltage[i]
			//n2pos = i
		}
	}

	//fmt.Printf("n2 joltage is %d at position %d\n", n2, n2pos)

	jolts, err := strconv.Atoi(fmt.Sprintf("%d%d", n1, n2))
	if err != nil {
		panic(err)
	}
	return jolts
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
