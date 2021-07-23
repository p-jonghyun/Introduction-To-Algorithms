package ch2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHornerRule(t *testing.T) {
	A := []int{1, 2, 3}
	naive := NaiveHornerRule(A, 2)
	fast := HornerRule(A, 2)

	assert.Equal(t, naive, 17)
	assert.Equal(t, fast, 17)
}
