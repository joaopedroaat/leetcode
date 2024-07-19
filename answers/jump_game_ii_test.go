package answers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{2, 3, 1, 1, 4}, expected: 2},
		{nums: []int{2, 3, 0, 1, 4}, expected: 2},
	}

	for _, test := range tests {
		actual := jump(test.nums)
		assert.Equalf(t, test.expected, actual, "Expected %d, got %d", test.expected, actual)
	}
}
