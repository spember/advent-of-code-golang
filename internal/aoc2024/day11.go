package aoc2024

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

// So naive! haha

func Day11Part1(input string, blinks int) int {
	initialStones := strings.Split(input, " ")

	var growStones = func(stones []string, remainingBlinks int) []string {
		// adjust a bit. build a massive initial array of stones, then grow them in place. avoid appending

		for i := 0; i < remainingBlinks; i++ {
			nextLineOfStones := make([]string, 0)
			for j := 0; j < len(stones); j++ {

				// for each stone, check left to right and apply one of 3 rules
				// rule 1, if 0, become 1
				if stones[j] == "0" {
					nextLineOfStones = append(nextLineOfStones, "1")
					continue
				}
				// rule 2
				if len(stones[j])%2 == 0 {
					left := aocutils.IntOrPanic(stones[j][:len(stones[j])/2])
					right := aocutils.IntOrPanic(stones[j][len(stones[j])/2:])
					nextLineOfStones = append(nextLineOfStones, fmt.Sprintf("%d", left), fmt.Sprintf("%d", right))
					continue
				}

				//rule 3 -> multiply by 2024
				nextLineOfStones = append(nextLineOfStones, fmt.Sprintf("%d", aocutils.IntOrPanic(stones[j])*2024))
			}
			stones = nextLineOfStones

		}

		return stones
	}

	finalStones := make([]string, 0)
	//var nonZero = 0
	// run for a capped number  of blinks, then restart with the new stones?
	var wg sync.WaitGroup
	for _, stone := range initialStones {
		wg.Add(1)
		calc := func() []string {
			defer wg.Done()
			return growStones([]string{stone}, blinks)
		}()
		finalStones = append(finalStones, calc...)

	}

	return len(finalStones)

	// the issue appears to be that the appending / growing of the line is too slow after 25 blinks
}
