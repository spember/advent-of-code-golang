package aoc202405

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

type Rule struct {
	PageNum int
	Before  []int
	After   []int
}

func NewRule(pageNum int) *Rule {
	return &Rule{
		PageNum: pageNum,
		Before:  []int{},
		After:   []int{},
	}
}

func Part1(input []byte) int {
	chunks := aocutils.ParseChunks(input)
	rules := createRules(chunks[0])
	validLines, _ := separateValidInstructions(rules, chunks[1])

	var total = 0
	for _, line := range validLines {
		if len(line) > 0 {
			total += line[len(line)/2]
		}
	}

	return total
}

func Part2(input []byte) int {
	chunks := aocutils.ParseChunks(input)
	rules := createRules(chunks[0])
	_, invalidLines := separateValidInstructions(rules, chunks[1])

	/*
			For each invalid line, we need to re-order the pages so that they are in a valid order.
		So, basically, we can use each page's 'before' rule to know which pages it should go after. we find the position
		of the right-most page in the NewOrder from the page's 'before' rules, and insert after that
	*/
	var middleSum = 0
	for _, line := range invalidLines {
		var newOrdering = make([]int, 0)
		for i, page := range line {
			if i == 0 {
				newOrdering = append(newOrdering, page)
				continue
			}
			//fmt.Printf("Page %d needs to come after %v and before %v\n", page, rules[page].Before, rules[page].After)
			// in the new order, find the right-most page that this page needs to come after
			var insertIndex = -1
			leftRules := rules[page].Before
			if len(leftRules) == 0 {
				//no left rules means this comes first
				newOrdering = append([]int{page}, newOrdering...)
				continue
			}
			if len(rules[page].After) == 0 {
				//no right rules means this comes last
				newOrdering = append(newOrdering, page)
				continue
			}
			// otherwise, find the right-most page in the newOrdering that this page needs to come after
			for p, newPage := range newOrdering {
				if slices.Contains(leftRules, newPage) {
					insertIndex = p
				}
			}

			//fmt.Printf("I think I need to insert %d at index %d\n ", page, insertIndex+1)
			newOrdering = slices.Insert(newOrdering, insertIndex+1, page)

		}
		//fmt.Println("New order is now", newOrdering)
		middleSum += newOrdering[len(newOrdering)/2]
	}

	return middleSum
}

func separateValidInstructions(rules map[int]*Rule, input []string) ([][]int, [][]int) {
	validLines := make([][]int, 0)
	invalidLines := make([][]int, 0)

	for _, pageLine := range input {
		pages := aocutils.Map(strings.Split(pageLine, ","), aocutils.IntOrPanic)
		// for each page, check that all pages after are in that page's after rules, and that all pages before are in the before rules
		var isValid = true
		for i, page := range pages {
			for _, afterPage := range pages[i+1:] {
				if !isAfter(page, afterPage, rules) {
					isValid = false
					break
				}
			}
			for _, beforePage := range pages[:i] {
				if !isBefore(page, beforePage, rules) {
					isValid = false
					break
				}
			}
		}
		//fmt.Printf("page line %d: %s is valid: %v\n", lineI, pageLine, isValid)
		if isValid {
			validLines = append(validLines, pages)
		} else {
			invalidLines = append(invalidLines, pages)
		}
	}
	return validLines, invalidLines
}

func isAfter(page, afterPage int, rules map[int]*Rule) bool {
	rule, ok := rules[page]
	if !ok {
		panic(fmt.Sprintf("No rule for page %d", page))
	}
	for _, after := range rule.After {
		if after == afterPage {
			return true
		}
	}
	return false
}

func isBefore(page, beforePage int, rules map[int]*Rule) bool {
	rule, ok := rules[page]
	if !ok {
		panic(fmt.Sprintf("No rule for page %d", page))
	}
	for _, before := range rule.Before {
		if before == beforePage {
			return true
		}
	}
	return false
}

func createRules(lines []string) map[int]*Rule {
	rules := make(map[int]*Rule)
	for _, chunk := range lines {
		pages := aocutils.Map(strings.Split(chunk, "|"), aocutils.IntOrPanic)
		
		var beforeRule *Rule
		beforeRule, ok := rules[pages[0]]
		if !ok {
			rules[pages[0]] = NewRule(pages[0])
			beforeRule = rules[pages[0]]
		}

		var afterRule *Rule
		afterRule, ok = rules[pages[1]]
		if !ok {
			rules[pages[1]] = NewRule(pages[1])
			afterRule = rules[pages[1]]
		}

		beforeRule.After = append(beforeRule.After, pages[1])
		afterRule.Before = append(afterRule.Before, pages[0])

	}
	return rules
}
