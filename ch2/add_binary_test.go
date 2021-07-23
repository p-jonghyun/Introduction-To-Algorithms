package ch2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBinary(t *testing.T) {

	t.Run("empty slices as input will give 0", func(t *testing.T) {
		C := AddBinary([]int{}, []int{})
		assert.Equal(t, len(C), 1)
		assert.Equal(t, C, []int{0})
	})

	t.Run("should work", func(t *testing.T) {
		A := []int{1, 1, 0}
		B := []int{0, 1, 0}
		C := AddBinary(A, B)
		assert.Equal(t, len(C), 4)
		assert.Equal(t, C, []int{1, 0, 0, 0})
	})
}
