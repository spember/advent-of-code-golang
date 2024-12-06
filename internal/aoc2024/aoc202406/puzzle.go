package aoc202406

import (
	"fmt"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

const free = "."
const obstacle = "#"

type Direction string

const (
	Up    Direction = "^"
	Right Direction = ">"
	Down  Direction = "v"
	Left  Direction = "<"
)

type Guard struct {
	// guard starts facing up
	Position     [2]int
	DirectionRow int
	DirectionCol int
	Direction    Direction
}

func (g *Guard) TurnRight() {
	switch g.Direction {
	case Up:
		g.Direction = Right
		g.DirectionRow = 0
		g.DirectionCol = 1
	case Right:
		g.Direction = Down
		g.DirectionRow = 1
		g.DirectionCol = 0
	case Down:
		g.Direction = Left
		g.DirectionRow = 0
		g.DirectionCol = -1
	case Left:
		g.Direction = Up
		g.DirectionRow = -1
		g.DirectionCol = 0
	}
	//fmt.Println("Guard turned right, now facing ", g.Direction)
}

func (g *Guard) TurnLeft() {
	switch g.Direction {
	case Up:
		g.Direction = Left
		g.DirectionRow = 0
		g.DirectionCol = -1
	case Right:
		g.Direction = Up
		g.DirectionRow = -1
		g.DirectionCol = 0
	case Down:
		g.Direction = Right
		g.DirectionRow = 0
		g.DirectionCol = 1
	case Left:
		g.Direction = Down
		g.DirectionRow = 1
		g.DirectionCol = 0
	}
	//fmt.Println("Guard turned left, now facing ", g.Direction)
}

func (g *Guard) WhatIsRight() Direction {
	switch g.Direction {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}
	panic("Invalid direction")
}

func (g *Guard) CheckInFront(grid [][]string) bool {
	// check if there is an obstacle in front of the guard
	x := g.Position[0] + g.DirectionRow
	y := g.Position[1] + g.DirectionCol
	return grid[x][y] == obstacle
}

func (g *Guard) CheckAtExit(grid [][]string) bool {
	// check if the guard is at the exit
	fx := g.Position[0] + g.DirectionRow
	fy := g.Position[1] + g.DirectionCol
	return g.IsOffGrid(grid, fx, fy)
}

func (g *Guard) IsOffGrid(grid [][]string, x, y int) bool {
	return x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0])
}

func (g *Guard) StepForward() {
	g.Position[0] += g.DirectionRow
	g.Position[1] += g.DirectionCol
}

func (g *Guard) StepBackward() {
	g.Position[0] -= g.DirectionRow
	g.Position[1] -= g.DirectionCol
}

func NewGuard(x, y int) *Guard {
	// guard starts facing up
	return &Guard{
		Position:     [2]int{x, y},
		DirectionRow: -1,
		DirectionCol: 0,
		Direction:    Up,
	}
}

func Part1(input []byte) int {
	grid := aocutils.ParseToStringGrid(input)
	// find guard
	var guard *Guard
	for i, row := range grid {
		for j, char := range row {
			if char != free && char != obstacle {
				fmt.Println("Found guard at ", i, j, " with char ", char)
				guard = NewGuard(i, j)
			}
		}
	}
	if guard == nil {
		panic("No guard found")
	}

	// follow the path until we reach the end
	var steps = 0
	var places = make(map[[2]int]bool)
	var placesWithDirection = make(map[[3]int]bool)
	places[guard.Position] = true
	placesWithDirection[[3]int{guard.Position[0], guard.Position[1], int(guard.Direction[0])}] = true

	for {
		placesWithDirection[[3]int{guard.Position[0], guard.Position[1], int(guard.Direction[0])}] = true
		// check if we're at the end
		if guard.CheckAtExit(grid) {
			fmt.Println("Guard is about to go off grid at ", guard.Position)
			steps++
			places[guard.Position] = true
			placesWithDirection[[3]int{guard.Position[0], guard.Position[1], int(guard.Direction[0])}] = true
			break
		}

		if guard.CheckInFront(grid) {
			guard.TurnRight()

		}
		// move forward
		guard.StepForward()
		steps++
		places[guard.Position] = true

		// turn right and step forward, if new position is in places With Direction, yay! we found a loop
		// then step backward and turn left
		guard.TurnRight()

		_, b := placesWithDirection[[3]int{guard.Position[0], guard.Position[1], int(guard.Direction[0])}]
		guard.TurnLeft()
		if b {
			fmt.Println("Found a loop at ", guard.Position[0]+guard.DirectionRow, guard.Position[1]+guard.DirectionCol)
		}

	}
	return len(places)
}
