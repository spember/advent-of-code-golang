package aoc2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Part1(t *testing.T) {
	d1 := Day1{}
	want := 0
	got := d1.Part1(d1SampleInput)
	assert.Equal(t, want, got)
}
