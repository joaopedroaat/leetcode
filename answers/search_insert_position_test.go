package answers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect int
	}{
		{nums: []int{1, 3, 5, 6}, target: 5, expect: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, expect: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, expect: 4},
	}

	for _, test := range tests {
		actual := searchInsert(test.nums, test.target)
		assert.Equalf(t, test.expect, actual, "Expected %d, got %d", test.expect, actual)
	}
}
