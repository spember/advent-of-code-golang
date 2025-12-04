package aocutils

// FindNextDiagonalSteps returns the next steps in a grid around the target position x,y
func FindNextDiagonalSteps[T any](grid [][]T, x, y int) [][]int {
	steps := [][]int{
		{x - 1, y - 1},
		{x - 1, y + 1},
		{x - 1, y},
		{x + 1, y - 1},
		{x + 1, y + 1},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	}

	return steps
}

func FindNextDiagonalStepsValid[T any](grid [][]T, x, y int) [][]int {
	validSteps := make([][]int, 0)
	steps := FindNextDiagonalSteps(grid, x, y)
	for _, step := range steps {
		if IsOnGrid(grid, step[0], step[1]) {
			validSteps = append(validSteps, step)
		}
	}
	return validSteps
}

func FindNextOrthogonalSteps[T any](grid [][]T, row, cell int) [][2]int {
	steps := make([][2]int, 0, 4)

	if IsOnGrid(grid, row-1, cell) {
		steps = append(steps, [2]int{row - 1, cell})
	}
	if IsOnGrid(grid, row+1, cell) {
		steps = append(steps, [2]int{row + 1, cell})
	}
	if IsOnGrid(grid, row, cell-1) {
		steps = append(steps, [2]int{row, cell - 1})
	}
	if IsOnGrid(grid, row, cell+1) {
		steps = append(steps, [2]int{row, cell + 1})
	}

	return steps
}

func IsOnGrid[T any](grid [][]T, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
}
