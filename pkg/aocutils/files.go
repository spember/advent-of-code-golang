package aocutils

import (
	"strings"
)

func ParseLines(fileData []byte) []string {
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

func ParseToStringGrid(fileData []byte) [][]string {
	lines := ParseLines(fileData)
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}
	return grid
}

func ParseAsRuneGrid(fileData []byte) [][]rune {
	lines := ParseLines(fileData)
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func ParseChunks(fileData []byte) [][]string {
	// convert to string
	lines := ParseLines(fileData)
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
