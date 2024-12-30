package aoc2024

import (
	"fmt"
	"math"
	"strings"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

func Day13Part1(input []byte, isOff bool) int {

	chunks := aocutils.ParseChunks(input)
	// each chunk contains two button lines and one prize line
	/*
		Button A: X+94, Y+34
		Button B: X+22, Y+67
		Prize: X=8400, Y=5400
	*/
	var totalTokens = 0
	for i, chunk := range chunks {
		aX, aY := parseButton(chunk[0], "+")
		bX, bY := parseButton(chunk[1], "+")
		prizeX, prizeY := parseButton(chunk[2], "=")
		if isOff {
			prizeX += 10000000000000
			prizeY += 10000000000000
		}

		// I had to research linear algebra for this. Here's a sample explanation on solving two variable equations

		/*

				You isolate for one variable then equate them. They're basically equations of the from y = mx + c, i.e 2 straight
				lines, what we want is an intersecting point among these 2 lines, so an (x,y) which satisfies both.

				If you have say ax + by = c and dx + ey = f then convert them to this by simply rearranging:

				y = (c - ax) / b

				y = (f - dx) / e

				then since y is same on both you get (c - ax) / b = (f - dx) / e, cross multiplying and rearranging you get

				x = (ec - bf)/(ea - bd)

			b=(py*ax-px*ay)/(by*ax-bx*ay) a=(px-b*bx)/ax
		*/
		b := float64(prizeY*aX-prizeX*aY) / float64(bY*aX-bX*aY)
		a := (float64(prizeX) - b*float64(bX)) / float64(aX)

		if a == math.Trunc(a) && b == math.Trunc(b) {
			fmt.Println("Test ", i+1, ": You win! (", a, ",", b, ")")
			totalTokens += (int(a) * 3) + (int(b) * 1)
		}

	}

	return totalTokens
}

func parseButton(line string, delim string) (x, y int) {
	// Button A: X+94, Y+34
	details := strings.Split(line, ": ")[1]
	xy := strings.Split(details, ", ")
	x = aocutils.IntOrPanic(strings.Replace(xy[0], "X"+delim, "", 1))
	y = aocutils.IntOrPanic(strings.Replace(xy[1], "Y"+delim, "", 1))
	return
}
