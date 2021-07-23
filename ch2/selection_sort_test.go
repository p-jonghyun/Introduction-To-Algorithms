package ch2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {

	t.Run("test for empty slice", func(t *testing.T) {
		var A []int
		SelectionSort(A)
		assert.Equal(t, 0, len(A))
	})

	t.Run("should sort", func(t *testing.T) {
		A := []int{5, 2, 3, 8, 1}
		SelectionSort(A)
		assert.Equal(t, len(A), 5)
		assert.Equal(t, A, []int{1, 2, 3, 5, 8})
	})

}