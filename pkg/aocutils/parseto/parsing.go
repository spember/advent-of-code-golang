package parseto

import (
	"iter"
	"strings"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

func Lines(fileData []byte) []string {
	// convert to string
	fileString := string(fileData)
	// split by newline
	lines := strings.Split(fileString, "\n")
	// remove last line if empty
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func LineSeqWithIndex(fileData []byte) iter.Seq2[int, string] {
	lines := Lines(fileData)
	return func(yield func(int, string) bool) {
		for p, line := range lines {
			if !yield(p, line) {
				break
			}
		}
	}
}

func LineSeq(fileData []byte) iter.Seq[string] {
	lines := Lines(fileData)
	return func(yield func(string) bool) {
		for _, line := range lines {
			if !yield(line) {
				break
			}
		}
	}
}

func Chunks(fileData []byte) [][]string {
	// convert to string
	lines := Lines(fileData)
	// split by empty line
	chunks := make([][]string, 0)
	chunk := make([]string, 0)
	for _, line := range lines {
		if line == "" || line == "\n" {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
		} else {
			chunk = append(chunk, line)
		}
	}
	// make sure to get the last chunk
	chunks = append(chunks, chunk)
	return chunks
}

func Grid[A any](fileData []byte, lineMapper func(string) []A) [][]A {
	lines := Lines(fileData)
	grid := make([][]A, len(lines))
	for i, line := range lines {
		grid[i] = lineMapper(line)
	}
	return grid
}

func StringGrid(fileData []byte, lineDelimiter string) [][]string {
	return Grid[string](fileData, func(s string) []string {
		return strings.Split(s, lineDelimiter)
	})
}

func IntGrid(fileData []byte) [][]int {
	return Grid[int](fileData, func(s string) []int {
		return aocutils.Map(strings.Split(s, ""), aocutils.IntOrPanic)
	})
}

func RuneGrid(fileData []byte, lineDelimiter string) [][]rune {
	return Grid[rune](fileData, func(s string) []rune {
		return []rune(s)
	})
}

/*
Mapper funcs
*/
