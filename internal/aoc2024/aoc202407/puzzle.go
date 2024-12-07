package aoc202407

import (
	"fmt"
	"strings"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

func Part1(input []byte) int {
	total := 0
	for _, line := range aocutils.ParseLines(input) {
		chunks := strings.Split(line, ": ")
		left := aocutils.IntOrPanic(chunks[0])
		right := aocutils.Map(strings.Split(chunks[1], " "), aocutils.IntOrPanic)
		discovered := discover(left, right[0], right[1:])
		if discovered > 0 {
			total += left
		}
	}
	return total
}

func discover(target int, current int, remaining []int) int {
	if len(remaining) == 0 {
		if current == target {
			return target
		} else {
			return 0
		}
	}
	withPlus := discover(target, current+remaining[0], remaining[1:])
	withTimes := discover(target, current*remaining[0], remaining[1:])
	return withPlus + withTimes
}

var operands = []string{"+", "*", "||"}
var basicOperands = []string{"+", "*"}

type Equations struct {
	Target  int
	Holding [][]string
}

func Part2(input []byte) int64 {

	//
	var total int64 = 0
	for _, line := range aocutils.ParseLines(input) {
		chunks := strings.Split(line, ": ")
		left := aocutils.IntOrPanic(chunks[0])
		right := aocutils.Map(strings.Split(chunks[1], " "), aocutils.IntOrPanic)
		// for each right, we need to build a set of equations, and then evaluate each one
		// recursively build equations, for each possible operand
		var equations = &Equations{Target: left, Holding: [][]string{}}
		buildEquation(equations, []string{fmt.Sprintf("%d", right[0])}, right[1:])

		for _, equation := range equations.Holding {
			result := solveEquation(equation)
			if result == left {
				fmt.Println("Found equation: ", equation, " = ", left)
				//fmt.Println("Result: ", result)
				total += int64(left)
				break
			}
		}

	}

	return total
}

func buildEquation(capturedEquations *Equations, currentEquation []string, remaining []int) {
	if len(remaining) == 0 {
		capturedEquations.Holding = append(capturedEquations.Holding, currentEquation)
		return
	}
	for _, op := range operands {
		buildEquation(capturedEquations, append(currentEquation, op, fmt.Sprintf("%d", remaining[0])), remaining[1:])
	}
	return
}

func solveEquation(equation []string) int {
	// for each operand, find the first one, and then evaluate the two numbers around it
	// then replace the three elements with the result
	// repeat until no operands are left

	var currentSum = aocutils.IntOrPanic(equation[0])
	var pos = 1
	for {
		if pos >= len(equation) {
			break
		}
		operand := equation[pos]
		nextValue := aocutils.IntOrPanic(equation[pos+1])

		switch operand {
		case "+":
			currentSum += nextValue
		case "*":
			currentSum *= nextValue
		case "||":
			currentSum = aocutils.IntOrPanic(fmt.Sprintf("%d%d", currentSum, nextValue))
		default:
			panic("Invalid operand")
		}

		pos += 2
	}

	return currentSum
}
