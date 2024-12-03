package aoc202403

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

var validInstruction = regexp.MustCompile(`mul\(\d+,\d+\)`)
var do = regexp.MustCompile(`do\(\)`)
var doNot = regexp.MustCompile(`don't\(\)`)

func Part1(input []byte) int64 {
	instructions := make([]string, 0)
	for _, line := range aocutils.ParseLines(input) {

		found := validInstruction.FindAll([]byte(line), -1)
		if found == nil {
			panic("Appears to be an invalid instruction")
		}
		for _, f := range found {
			instructions = append(instructions, string(f))
		}
	}

	var result int64
	for _, instruction := range instructions {
		result += convertInstructions(instruction)
	}
	return result
}

func Part2(input []byte) int64 {
	var total int64 = 0
	var allowed = true
	for _, line := range aocutils.ParseLines(input) {
		//fmt.Println("Going into line with allowed: ", allowed)
		lineBytes := []byte(line)

		// FindAllIndex returns a slice of pairs of indices where the regex matches AND where it ends, thus we can get
		// the size of the 'jump' to make each iteration, reducing the number of iterations.
		doSegments := do.FindAllIndex(lineBytes, -1)
		doNotSegments := doNot.FindAllIndex(lineBytes, -1)
		instructionSegments := validInstruction.FindAllIndex(lineBytes, -1)
		fmt.Println(doSegments, doNotSegments, instructionSegments)
		// for each rune in line, check if it's a DO or a Don't, switch on and off accordingly
		// if it's an instruction, either ignore or convert it

		for i := 0; i < len(line); {
			var increase = 1
			if len(doSegments) > 0 && i == doSegments[0][0] {
				//fmt.Println("Found DO at ", doSegments[0][0])
				allowed = true
				increase = doSegments[0][1] - doSegments[0][0]
				doSegments = doSegments[1:]

			}
			if len(doNotSegments) > 0 && i == doNotSegments[0][0] {
				//fmt.Println("Found DON'T at ", doNotSegments[0][0])
				allowed = false
				increase = doNotSegments[0][1] - doNotSegments[0][0]
				doNotSegments = doNotSegments[1:]

			}

			if len(instructionSegments) > 0 && i == instructionSegments[0][0] {
				//fmt.Println("Found instruction at ", instructionSegments[0][0])
				if allowed {
					total += convertInstructions(line[i:instructionSegments[0][1]])
				} else {
					// ignore!
					//fmt.Println("Ignoring instruction at ", instructionSegments[0][0])
				}
				increase = instructionSegments[0][1] - instructionSegments[0][0]
				instructionSegments = instructionSegments[1:]
			}

			i += increase
		}
	}
	return total
}

func convertInstructions(instruction string) int64 {
	pairs := strings.Split(instruction[4:len(instruction)-1], ",")
	step := aocutils.IntOrPanic(pairs[0]) * aocutils.IntOrPanic(pairs[1])
	//fmt.Println(step)
	return int64(step)
}
