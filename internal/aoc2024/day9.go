package aoc2024

import (
	"fmt"
	"strings"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

//const freeSpace = "."

type FileBuffer struct {
	Buffer   []string
	Shuffled map[string]bool
}

func (f *FileBuffer) calcChecksum() int {
	var checksum = 0
	for i := 0; i < len(f.Buffer); i++ {
		if f.Buffer[i] == freeSpace {
			continue
		}
		checksum += i * aocutils.IntOrPanic(f.Buffer[i])
	}

	return checksum
}

func (f *FileBuffer) FindLeftMostFreeSpace(maxPosition, size int) (int, bool) {
	// find the leftmost free contiguous space of size int. return the position and a bool indicating if it was found
	var scanChunk = func(start int) (int, bool) {
		var found = true
		for s := start; s < start+size; s++ {

			if s >= maxPosition || f.Buffer[s] != freeSpace {
				found = false
				break
			}
		}
		return start, found
	}

	var freePos = 0
	//var found = false
	for i := 0; i < maxPosition+size; {
		if f.Buffer[i] == freeSpace {
			// we have one... but is it enough?
			freePos = i
			p, ok := scanChunk(freePos)
			if ok {
				//fmt.Println("Found a set of spaces at ", p, size)
				return p, true
			}
			i += size
		} else {
			i++
		}
	}
	return 0, false
}

func (f *FileBuffer) ShuffleFile(fileId string, fPos, size int) {
	_, c := f.Shuffled[fileId]
	if c {
		//fmt.Printf("File %s already shuffled\n", fileId)
		return
	} else {
		f.Shuffled[fileId] = true
	}

	open, found := f.FindLeftMostFreeSpace(fPos, size)
	if !found {
		//fmt.Println("No free space found for file ", fileId)
		return
	}
	//fmt.Println("Shuffling file ", fileId, "from ", fPos, "to ", open)
	for i := 0; i < size; i++ {
		f.Buffer[open+i] = f.Buffer[fPos+i]
		f.Buffer[fPos+i] = freeSpace
	}
	//fmt.Println(f.Buffer)
}

func NewFileBuffer(values []int) *FileBuffer {
	totalInputSize := 0
	for _, v := range values {
		totalInputSize += v
	}
	buffer := make([]string, totalInputSize)
	fmt.Println("Mem allocated , size: ", totalInputSize)
	// seed buffer
	var idCounter = 0
	var bufferPos = 0
	for i := 0; i < len(values); {
		fileId := fmt.Sprintf("%d", idCounter)
		//fmt.Println("File is ", fileId)
		for f := 0; f < values[i]; f++ {
			buffer[bufferPos] = fileId
			bufferPos++
		}

		i++
		idCounter++
		if i >= len(values) {
			break
		}
		//fmt.Println("Empty is ", values[i])
		for e := 0; e < values[i]; e++ {
			buffer[bufferPos] = freeSpace
			bufferPos++
		}
		i++
	}
	return &FileBuffer{
		Buffer:   buffer,
		Shuffled: make(map[string]bool),
	}
}

func Day9Part1(input []byte) int {

	// find the total length of the input
	lines := aocutils.ParseLines(input)
	if len(lines) < 1 {
		fmt.Println("No lines in input")
		return 0
	}

	values := aocutils.Map(strings.Split(lines[0], ""), aocutils.IntOrPanic)
	fileBuffer := NewFileBuffer(values)

	// start shuffling

	var freePos = 0
	var filePos = len(fileBuffer.Buffer) - 1
	for {
		// find the next free space
		// find the rightmost file
		// and swap them
		// keep doing that until the freePos is greater than equal to filePos
		if freePos >= filePos {
			break
		}
		for fileBuffer.Buffer[freePos] != freeSpace {
			freePos++
		}
		for fileBuffer.Buffer[filePos] == freeSpace {
			filePos--
		}
		if freePos >= filePos {
			break
		}
		fileBuffer.Buffer[freePos], fileBuffer.Buffer[filePos] = fileBuffer.Buffer[filePos], fileBuffer.Buffer[freePos]
	}
	fmt.Println(fileBuffer.Buffer)

	return fileBuffer.calcChecksum()
}

func Day9Part2(input []byte) int {
	lines := aocutils.ParseLines(input)
	if len(lines) < 1 {
		fmt.Println("No lines in input")
		return 0
	}

	values := aocutils.Map(strings.Split(lines[0], ""), aocutils.IntOrPanic)
	fileBuffer := NewFileBuffer(values)

	var currentFileId = fileBuffer.Buffer[len(fileBuffer.Buffer)-1]
	var currentFileSize = 1
	var currentFilePos = 0
	for filePos := len(fileBuffer.Buffer) - 2; filePos > 0; filePos-- {

		if fileBuffer.Buffer[filePos] == freeSpace {
			//fmt.Println("Free space at ", filePos, "shuffle file!")
			fileBuffer.ShuffleFile(currentFileId, currentFilePos, currentFileSize)
			continue
		}
		if fileBuffer.Buffer[filePos] != currentFileId {
			//fmt.Println("New file found at ", filePos, "old file of id ", currentFileId, "size ", currentFileSize, "shuffle file!")
			// new file! have to move the old one!
			fileBuffer.ShuffleFile(currentFileId, currentFilePos, currentFileSize)
			currentFileId = fileBuffer.Buffer[filePos]
			currentFileSize = 1
			currentFilePos = filePos
		} else {
			currentFileSize++
			currentFilePos = filePos
			continue
		}
	}
	fmt.Println(fileBuffer.Buffer)

	return fileBuffer.calcChecksum()
}
