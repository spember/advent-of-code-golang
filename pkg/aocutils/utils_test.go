package aocutils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWindowExact(t *testing.T) {

	sample := "hello, there"

	var hitCount = 0
	WindowExact(sample, 3, 1, func(sub string, pos int) {
		fmt.Println(sub, pos)
		hitCount += 1
	})
	assert.Equal(t, 7, hitCount)

}
