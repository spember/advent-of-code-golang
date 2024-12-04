package aoc202404

import (
	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

/*
Aoc 2024 day 4 puzzle

*/

const xmasLiteral = "XMAS"

func Part1(input []byte) int {
	grid := aocutils.ParseToStringGrid(input)
	var total = 0
	for x, row := range grid {
		for y, char := range row {
			if char == string(xmasLiteral[0]) {
				total += countWords(grid, x, y)
			}
		}
	}
	return total
}

func countWords(grid [][]string, x, y int) int {
	var total = 0
	words := []string{
		extractWord(grid, []int{x, y}, 0, 1),
		extractWord(grid, []int{x, y}, 0, -1),
		extractWord(grid, []int{x, y}, 1, 0),
		extractWord(grid, []int{x, y}, -1, 0),

		extractWord(grid, []int{x, y}, 1, 1),
		extractWord(grid, []int{x, y}, 1, -1),

		extractWord(grid, []int{x, y}, -1, 1),
		extractWord(grid, []int{x, y}, -1, -1),
	}

	for _, word := range words {
		//fmt.Printf("start %d,%d, Checking word %s\n", x, y, word)
		if word == xmasLiteral {
			//fmt.Println("Found word!")
			total++
		}
	}
	return total
}

func extractWord(grid [][]string, start []int, xInc, yInc int) string {
	var word = ""

	x := start[0]
	y := start[1]

	for i := 0; i < len(xmasLiteral); i++ {
		// if we're out of bounds, return the word
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
			break
		}
		word += grid[x][y]
		x += xInc
		y += yInc
	}
	return word
}

const center = "A"

const MAS = "MAS"
const SAM = "SAM"

func Part2(input []byte) int {
	// X-Mas. Looking for "MAS" , diagonals only
	// find an A, then extract -1,+1, 0,0, +1, -1, and -1, -1, 0,0, and +1, +1
	// if word 1 == MAS or SAM AND word 2 == MAS or SAM, return 1

	grid := aocutils.ParseToStringGrid(input)
	var total = 0
	for x, row := range grid {
		for y, char := range row {
			if char == center {
				if checkMas(grid, []int{x, y}) {
					//fmt.Println("Is MAS!")
					total++
				}
			}
		}
	}
	return total
}

func checkMas(grid [][]string, start []int) bool {
	x := start[0]
	y := start[1]

	//-1,+1, 0,0, +1, -1, and -1, -1, 0,0, and +1, +1
	words := []string{
		extractMas(grid, [][]int{
			{x - 1, y + 1},
			{x, y},
			{x + 1, y - 1},
		}),
		extractMas(grid, [][]int{
			{x - 1, y - 1},
			{x, y},
			{x + 1, y + 1},
		}),
	}

	if (words[0] == MAS || words[0] == SAM) && (words[1] == MAS || words[1] == SAM) {
		return true
	}
	return false
}

func extractMas(grid [][]string, sequence [][]int) string {
	var word = ""
	for _, pos := range sequence {
		x := pos[0]
		y := pos[1]
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
			break
		}
		word += grid[x][y]
	}
	return word
}
