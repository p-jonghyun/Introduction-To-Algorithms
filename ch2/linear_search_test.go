package ch2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearSearch(t *testing.T) {

	t.Run("should return index if exists", func(t *testing.T) {
		index := LinearSearch([]int{1, 2, 3}, 3)
		assert.Equal(t, index, 2)
	})

	t.Run("should return -1 if v doesn't exist", func(t *testing.T) {
		index := LinearSearch(nil, 3)
		assert.Equal(t, index, -1)
	})
}