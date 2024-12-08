package aoc2024

import (
	"fmt"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

const freeSpace = "."

func Day8Part1(input []byte) int {

	grid := aocutils.ParseToStringGrid(input)
	antenae := findAntenae(grid)

	fmt.Println(antenae)
	// for each antenae signal, we need to find the antinodes of each pair
	antinodes := make(map[[2]int]bool, 0)

	for signal, antenaePairs := range antenae {
		fmt.Println("Signal: ", signal)
		//fmt.Println("Antenae pairs: ", antenaePairs)

		for p, _ := range antenaePairs {
			var o = 0
			for o < len(antenaePairs) {
				if o == p {
					o++
					continue
				}
				fmt.Printf("Comparing %v and %v\n", antenaePairs[p], antenaePairs[o])
				// find the antinodes of the pair
				first := [2]int{
					antenaePairs[p][0] + (antenaePairs[p][0] - antenaePairs[o][0]),
					antenaePairs[p][1] + (antenaePairs[p][1] - antenaePairs[o][1]),
				}
				second := [2]int{
					antenaePairs[o][0] + (antenaePairs[o][0] - antenaePairs[p][0]),
					antenaePairs[o][1] + (antenaePairs[o][1] - antenaePairs[p][1]),
				}
				
				antinodes[first] = true
				antinodes[second] = true
				o++
			}

		}

	}

	var onboardCount = 0
	for anti, _ := range antinodes {
		// if anti on grid
		if aocutils.IsOnGrid(grid, anti[0], anti[1]) {
			onboardCount++
		}
	}

	return onboardCount
}

func Day8Part2(input []byte) int {
	grid := aocutils.ParseToStringGrid(input)
	antenae := findAntenae(grid)

	antinodes := make(map[[2]int]bool, 0)

	for signal, antenaePairs := range antenae {
		fmt.Println("Signal: ", signal)
		//fmt.Println("Antenae pairs: ", antenaePairs)

		for p, _ := range antenaePairs {
			var o = 0
			for o < len(antenaePairs) {
				if o == p {
					o++
					continue
				}

				// make new antinodes in both directions until off the board
				rowRate := antenaePairs[p][0] - antenaePairs[o][0]
				colRate := antenaePairs[p][1] - antenaePairs[o][1]

				anti := [2]int{
					antenaePairs[p][0] + rowRate,
					antenaePairs[p][1] + colRate,
				}
				for {
					if !aocutils.IsOnGrid(grid, anti[0], anti[1]) {
						break
					} else {
						antinodes[anti] = true
						anti = [2]int{
							anti[0] + rowRate,
							anti[1] + colRate,
						}
					}
				}

				anti = [2]int{
					antenaePairs[p][0] - rowRate,
					antenaePairs[p][1] - colRate,
				}
				for {
					if !aocutils.IsOnGrid(grid, anti[0], anti[1]) {
						break
					} else {
						antinodes[anti] = true
						anti = [2]int{
							anti[0] - rowRate,
							anti[1] - colRate,
						}
					}
				}
				o++
			}

		}

	}

	var onboardCount = 0
	for anti, _ := range antinodes {
		if aocutils.IsOnGrid(grid, anti[0], anti[1]) {
			onboardCount++
		}
	}
	return onboardCount
}

func findAntenae(grid [][]string) map[string][][]int {
	antenae := make(map[string][][]int)

	for rowP, line := range grid {
		fmt.Println(line)
		for colP, space := range line {
			if space == freeSpace {
				continue
			}
			antenae[space] = append(antenae[space], []int{rowP, colP})
		}
	}
	return antenae
}
