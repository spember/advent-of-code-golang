package aoc2023

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"unicode"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

//go:embed day01-sample.txt
var d1SampleInput []byte

//go:embed day01-sample2.txt
var d1Sample2Input []byte

type Day1 struct {
}

func (d *Day1) Part1(data []byte) int {

	scanForDigit := func(runes []rune) (rune, int) {
		for i, r := range runes {
			if unicode.IsDigit(r) {
				return r, i
			}
		}
		panic("We didn't find one")
	}

	extract := func(line string) (int, error) {
		runes := []rune(line)

		lvalue, _ := scanForDigit(runes)
		slices.Reverse(runes)
		rvalue, _ := scanForDigit(runes)

		fmt.Printf("Found %s%s\n", string(lvalue), string(rvalue))

		return strconv.Atoi(fmt.Sprintf("%s%s", string(lvalue), string(rvalue)))
	}

	var total = 0

	for i, l := range aocutils.ParseLines(data) {
		fmt.Printf("%d: %s\n", i, l)
		value, err := extract(l)
		if err != nil {
			panic("Could not find a value!")
		}
		total += value
	}
	fmt.Println("Done!")
	return total
}

type NumberLocation struct {
	Position int
	Value    int
}

var words *map[string]int = &map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func (d *Day1) Part2(data []byte) int {
	// for each line, find all numbers and their indexes

	search := func(line string) int {

		return 0
	}

	var sum = 0
	for _, l := range aocutils.ParseLines(data) {
		sum += search(l)
	}
	return 0
}
