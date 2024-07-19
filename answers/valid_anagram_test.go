package answers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{s: "anagram", t: "nagaram", expected: true},
		{s: "rat", t: "car", expected: false},
	}
	for _, test := range tests {
		actual := isAnagram(test.s, test.t)
		assert.Equalf(t, test.expected, actual, "Expected %t, got %t", test.expected, actual)
	}
}
