package aoc2023

import (
	_ "embed"
	"fmt"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

//go:embed day01-sample.txt
var d1SampleInput []byte

type Day1 struct {
}

func (d *Day1) Part1(data []byte) int {

	//extract := func(line string) (int) {
	//	var left = -1
	//	var right = -1
	//	var lp = 0
	//	var rp = len(line)-1
	//
	//	for {
	//
	//		if lp >= rp {
	//			panic("threshold crossed")
	//		}
	//		if left >= 0 && right >= 0 {
	//			break
	//		}
	//	}
	//
	//}

	for i, l := range aocutils.ParseLines(data) {
		fmt.Printf("%d: %s\n", i, l)

		runes := []rune(l)
		for _, c := range runes {
			fmt.Println(c)
		}
	}
	fmt.Println("Done!")
	return 0
}
