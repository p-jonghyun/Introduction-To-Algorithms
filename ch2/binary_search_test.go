package ch2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterativeBinarySearch(t *testing.T) {
	t.Run("should return index if exists", func(t *testing.T) {
		index := IterativeBinarySearch([]int{1, 2, 3}, 3, 0, 2)
		assert.Equal(t, index, 2)
	})
	t.Run("return -1 if not exist", func(t *testing.T) {
		index := IterativeBinarySearch([]int{1, 5, 6}, 2, 0, 2)
		assert.Equal(t, index, -1)
	})
}

func TestRecursiveBinarySearch(t *testing.T) {
	t.Run("should return index if exists", func(t *testing.T) {
		index := RecursiveBinarySearch([]int{1, 2, 3}, 3, 0, 2)
		assert.Equal(t, index, 2)
	})
	t.Run("return -1 if not exist", func(t *testing.T) {
		index := RecursiveBinarySearch([]int{1, 5, 6}, 2, 0, 2)
		assert.Equal(t, index, -1)
	})
}