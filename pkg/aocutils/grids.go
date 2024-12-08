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

func IsOnGrid[T any](grid [][]T, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
}
