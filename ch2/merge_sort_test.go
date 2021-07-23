package ch2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {

	t.Run("test for empty slice", func(t *testing.T) {
		var A []int
		MergeSort(A, 0, 0)
		assert.Equal(t, 0, len(A))
	})

	t.Run("should sort", func(t *testing.T) {
		A := []int{5, 2, 3, 8, 1}
		MergeSort(A, 0, 4)
		assert.Equal(t, len(A), 5)
		assert.Equal(t, A, []int{1, 2, 3, 5, 8})
	})

}

func TestMerge(t *testing.T) {
	t.Run("should correctly merge", func(t *testing.T) {
		A := []int{3, 5, 1, 4}
		Merge(A, 0, 1, 3)
		fmt.Println(A)
	})
}
