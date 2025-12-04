package aoc2025

import (
	"strings"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
	"github.com/spember/advent-of-code-golang/pkg/aocutils/printer"
)

type PrintingDepartment struct {
}

func (p *PrintingDepartment) SolveP1(grid [][]string) int {
	return len(p.DiscoverMovableRolls(grid))
}

func (p *PrintingDepartment) SolveP2(grid [][]string) int {
	// iterate on discovery, replacing each roll with a dot
	var iterations = 1
	var rollsRemoved = 0

	for {
		rollPositions := p.DiscoverMovableRolls(grid)
		printer.F("Iter %d: Discovered rolls: %v\n", iterations, rollPositions)
		rollsRemoved += len(p.DiscoverMovableRolls(grid))

		// now actually remove them:
		for _, pos := range rollPositions {
			grid[pos[0]][pos[1]] = "."
		}

		if len(rollPositions) == 0 {
			printer.F("Iter %d: No rolls discovered, breaking\n", iterations)
			break
		}

		if iterations >= 1000 {
			printer.Ln("Breaking due to iteration cap hit")
			break
		}
		iterations++
	}

	return rollsRemoved
}

func (p *PrintingDepartment) DiscoverMovableRolls(grid [][]string) [][]int {
	//for _, row := range grid {
	//	for _, col := range row {
	//		fmt.Print(col, " ")
	//	}
	//	fmt.Println()
	//}
	var rollPositions = make([][]int, 0)
	for r, row := range grid {
		for c, _ := range row {
			rolls := p.CountSymbolsIn(grid, r, c, "@")
			if rolls < 4 && grid[r][c] == "@" {
				//printer.Ln("***There is a roll with less than 4 around it at ", r, ",", c)
				rollPositions = append(rollPositions, []int{r, c})
			}
		}
	}
	return rollPositions
}

func (p *PrintingDepartment) CountSymbolsIn(grid [][]string, r, c int, symbol string) int {
	nextSteps := aocutils.FindNextDiagonalStepsValid[string](grid, r, c)
	//printer.F("%d,%d: has next steps: %v\n", r, c, nextSteps)
	var count = 0
	for _, step := range nextSteps {
		if strings.Contains(grid[step[0]][step[1]], symbol) {
			count++
		}
	}
	return count
}
