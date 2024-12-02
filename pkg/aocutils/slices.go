package aocutils

import "strconv"

// WindowExact performs an 'exact' windowing over the string. No Partial windows allowed
func WindowExact(input string, size int, step int, receiver func(chunk string, position int)) {
	endpoint := len(input)
	position := 0
	for {
		// 5, 0 2
		if (position + size) > (endpoint - size) {
			return
		}
		receiver(input[position:(position+size)], position)
		position += step
	}
}

// IntOrPanic converts a string to an int or panics, only useful for AOC puzzles or other places where we fully expect an int
func IntOrPanic(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}

// SliceCount counts the number of times a value appears in a slice
func SliceCount(slice []int, value int) int {
	var count = 0
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}
