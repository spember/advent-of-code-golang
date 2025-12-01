package aoc2025

import (
	"fmt"
	"strconv"

	"github.com/spember/advent-of-code-golang/pkg/aocutils/parseto"
	"github.com/spember/advent-of-code-golang/pkg/aocutils/printer"
)

/*
Day 1 Secret Entrence

https://adventofcode.com/2025/day/1

You could follow the instructions, but your recent required official North Pole secret entrance security training seminar taught you that the safe is actually a decoy. The actual password is the number of times the dial is left pointing at 0 after any rotation in the sequence.
*/

type SecretEntrance struct {
}

func (s SecretEntrance) Part1(input []byte) int {
	lines := parseto.Lines(input)

	var currentPosition int = 50
	var zeroes int = 0

	for _, line := range lines {
		dir := string(line[0])
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch dir {
		case "L":
			//fmt.Println("Turning left by ", amount)
			currentPosition = s.turnDial(currentPosition - amount)
		case "R":

			currentPosition = s.turnDial(currentPosition + amount)
		}

		if currentPosition == 0 {
			zeroes++
		}
	}

	return zeroes
}

func (s SecretEntrance) turnDial(amount int) int {
	// if pos goes below zero, start at 99 again
	//fmt.Println("Proposed next position is ", amount)

	if amount < -99 {
		fmt.Println("huge Proposed next position is ", amount)
		amount = amount % 100
	}
	if amount > 99 {
		fmt.Println("huge Proposed next position is ", amount)
		amount = amount % 100
	}
	nextPos := amount

	if nextPos < 0 {
		nextPos = 100 + nextPos
		//fmt.Println("shifted from to, ", amount, nextPos)
	} else if nextPos >= 100 {
		nextPos = nextPos - 100
		//fmt.Println("shifted from to, ", amount, nextPos)
	}
	return nextPos

}

func (s SecretEntrance) Part2(lines []string) int {

	var currentPosition int = 50
	var stepClicks = 0
	var zeroes = 0

	for _, line := range lines {
		dir := string(line[0])
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch dir {
		case "L":
			amount = -(amount)
		}

		currentPosition, stepClicks = s.turnForClicksBrute(currentPosition, amount)
		printer.Ln("Setting new position to ", currentPosition, "with clicks = ", stepClicks)
		zeroes += stepClicks
	}

	return zeroes
}

func (s SecretEntrance) turnForClicksBrute(incomingPosition int, amount int) (int, int) {
	clicks := 0
	currentPosition := incomingPosition
	// I am double counting somewhere
	// if starting at 0, don't count when moving off. If ending at zero, don't  count the move off
	// however, we do want to count
	if amount > 0 {
		// go up
		printer.Ln("Going up by ", amount)
		for i := 0; i < amount; i++ {
			if currentPosition == 0 {
				printer.Ln("We're moving off of zero!")
				clicks++
			}
			currentPosition++
			if currentPosition >= 100 {
				currentPosition = 0
			}
		}
	} else {
		printer.Ln("Going down by ", amount, "from ", incomingPosition)
		for i := amount; i < 0; i++ {
			if currentPosition == 0 {
				printer.Ln("We're moving off of zero!")
				clicks++
			}
			currentPosition--
			if currentPosition < 0 {
				currentPosition = 99
			}
		}
	}

	if incomingPosition == 0 {
		clicks--
	}
	if currentPosition == 0 {
		clicks++
	}

	printer.Ln("Adjusted ", incomingPosition, "to be ", currentPosition)
	return currentPosition, clicks
}
