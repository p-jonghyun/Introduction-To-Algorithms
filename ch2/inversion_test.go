package ch2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInversions(t *testing.T) {
	inversions := Inversions([]int{2, 3, 8, 6, 1}, 0, 4)
	assert.Equal(t, 5, inversions)
}
