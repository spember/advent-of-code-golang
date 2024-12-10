package aoc2024

import (
	"slices"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

/*
I accidentally solved part 2 first, oops
*/

func Day10(input []byte) (int, int) {
	grid := aocutils.ParseToIntGrid(input)

	var total = 0
	var rating = 0
	for r, row := range grid {
		for c, cell := range row {
			if cell == 0 {
				score := plotTrail(grid, [2]int{-1, -1}, [2]int{r, c})

				prunedScore := make([][2]int, 0)
				for _, s := range score {
					if !slices.Contains(prunedScore, s) {
						prunedScore = append(prunedScore, s)
					}
				}
				total += len(prunedScore)
				rating += len(score)
			}
		}
	}

	return total, rating
}

func plotTrail(grid [][]int, previous [2]int, current [2]int) [][2]int {

	if grid[current[0]][current[1]] == 9 {
		// fantastic, we have reached the end
		//fmt.Println("Reached the end. Found 9 at ", current)
		return [][2]int{current}
	}

	orthogonalSteps := aocutils.FindNextOrthogonalSteps(grid, current[0], current[1])
	tops := make([][2]int, 0)
	for _, step := range orthogonalSteps {
		if step == previous {
			continue
		}
		if grid[step[0]][step[1]] == grid[current[0]][current[1]]+1 {
			tops = append(tops, plotTrail(grid, current, step)...)
		}
	}
	return tops

}
