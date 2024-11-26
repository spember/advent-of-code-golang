package aocutils

import "strings"

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

func ParseChunks(fileData []byte) [][]string {
	// convert to string
	lines := ParseLines(fileData)
	// split by empty line
	chunks := make([][]string, 0)
	chunk := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
		} else {
			chunk = append(chunk, line)
		}
	}
	return chunks
}
