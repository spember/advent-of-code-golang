package aoc2025

import (
	"fmt"
	"iter"
	"strconv"
	"strings"

	"github.com/spember/advent-of-code-golang/pkg/aocutils/printer"
)

type GiftShop struct{}

func (g GiftShop) Solve(idChunks iter.Seq[[]int], useBasic bool) int {
	var invalidIds []int = make([]int, 0)
	for c := range idChunks {
		printer.F("%d -> %d\n", c[0], c[1])
		if c[0] > c[1] {
			panic("c[0] > c[1]" + fmt.Sprintf("%d", c[0]) + ", " + fmt.Sprintf("%d", c[1]))
		}
		for n := c[0]; n <= c[1]; n++ {
			if useBasic {
				if ok := g.BasicIsIdValid(n); !ok {
					invalidIds = append(invalidIds, n)
				}
			} else {
				if ok := g.RepeatingIsIdValid(n); !ok {
					invalidIds = append(invalidIds, n)
				}
			}

		}
		printer.F("After sequence, invalid ids are %v\n ", invalidIds)
	}

	var sum int
	for _, i := range invalidIds {
		sum += i
	}

	printer.Ln(invalidIds)
	printer.Ln(sum)
	return sum

}

func (g GiftShop) BasicIsIdValid(id int) bool {
	repr := strconv.Itoa(id)
	runes := []rune(repr)

	// can only handle even length
	if len(runes)%2 != 0 {
		return true
	}
	mid := len(runes) / 2

	firstHalf := string(runes[:mid])
	secondHalf := string(runes[mid:])

	if firstHalf == secondHalf {
		printer.Ln("Invalid id! " + repr)
		return false
	}
	return true
}

func (g GiftShop) RepeatingIsIdValid(id int) bool {
	repr := strconv.Itoa(id)

	mid := len(repr) / 2

	// go up from 1 to mid and count repeating
	// however,

	for i := 1; i <= mid; i++ {
		segment := repr[:i]
		withRemoved := strings.ReplaceAll(repr, segment, "")
		printer.F("Looking for %s in %s (mid is %d) With Removed is %s\n", repr[:i], repr, mid, withRemoved)
		if strings.Count(repr, segment) >= 2 && withRemoved == "" {
			return false
		}
	}
	return true
}

func (g GiftShop) ParseLine(line string) iter.Seq[[]int] {
	segments := strings.Split(line, ",")
	return func(yield func([]int) bool) {
		for _, segment := range segments {
			parts := strings.Split(segment, "-")
			left, _ := strconv.Atoi(parts[0])
			right, _ := strconv.Atoi(parts[1])
			yield([]int{left, right})
		}
	}
}
