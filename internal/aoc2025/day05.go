package aoc2025

import (
	"fmt"
	"strconv"

	"github.com/spember/advent-of-code-golang/pkg/aocutils/printer"
)

type Cafeteria struct {
}

func (c *Cafeteria) SolveP1(idRanges []string, availableIds []string) int {
	validRanges := make([][]int, len(idRanges))
	for p, id := range idRanges {
		validRanges[p] = c.ConvertToRange(id)
		//fmt.Println(validRanges[p])
	}
	//fmt.Println()

	// ok, now check which spoiledIds are actually spoiled
	var validCount int = 0
	for _, id := range availableIds {
		ingredientId, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		for _, validRange := range validRanges {
			if c.CheckIfInRange(ingredientId, validRange) {
				validCount++
				break
			}
		}

	}

	return validCount
}

func (c *Cafeteria) ConvertToRange(input string) []int {
	var result []int = make([]int, 2)
	n, err := fmt.Sscanf(input, "%d-%d", &result[0], &result[1])
	if err != nil || n != 2 {
		panic("Failed to parse range: " + input)
	}
	return result
}

func (c *Cafeteria) CheckIfInRange(id int, validRange []int) bool {
	return id >= validRange[0] && id <= validRange[1]
}

func (c *Cafeteria) SolveP2(idRanges []string) int {
	// find the total covered range of all sub ranges
	validRanges := make([][]int, len(idRanges))

	for p, id := range idRanges {
		validRanges[p] = c.ConvertToRange(id)
		//printer.Ln("Range:", validRanges[p])
	}
	//fmt.Println()

	// find the overlapping ranges and merge to create larger ranges
	// now iterate over the ranges, creating new ones or merging with existing
	// if no ranges were merged, break
	var itercount = 0
	var ok bool = false
	for {
		validRanges, ok = c.MergeRangeSinglePass(validRanges)
		if !ok {
			printer.F("Printer: No merges on iteration %d, breaking\n", itercount)
			break
		}
		//for _, r := range validRanges {
		//	printer.Ln("Merged Range:", r)
		//}

		itercount++
		if itercount > 100 {
			fmt.Println("Breaking due to iteration count")
			break
		}
	}

	// count ranges
	var coveredCount int = 0
	for _, r := range validRanges {
		coveredCount += (r[1] - r[0] + 1)
	}
	return coveredCount
}

func (c *Cafeteria) MergeRangeSinglePass(ranges [][]int) ([][]int, bool) {
	mergedRanges := make([][]int, 0)
	mergedAny := false

	//for _, currentRange := range ranges {
	//	merged := false
	//	for i, existingRange := range mergedRanges {
	//		if c.RangesOverlap(existingRange, currentRange) {

	// for each range in ranges, check to see if it overlaps with any in mergedRanges
	// if not, add it to mergedRanges
	// if so, merge it and mark mergedAny = true
	for _, currentRange := range ranges {
		merged := false

		for _, existingRange := range mergedRanges {
			// if overlap
			if c.RangesOverlap(currentRange, existingRange) {
				existingRange[0] = min(existingRange[0], currentRange[0])
				existingRange[1] = max(existingRange[1], currentRange[1])
				merged = true
				break
			}
		}
		if merged {
			mergedAny = true
		} else {
			mergedRanges = append(mergedRanges, currentRange)
		}
	}

	return mergedRanges, mergedAny
}

func (c *Cafeteria) RangesOverlap(rangeA, rangeB []int) bool {
	// two ranges overlap if the start of one is less than or equal to the end of the other
	return rangeA[0] <= rangeB[1] && rangeB[0] <= rangeA[1]
}
