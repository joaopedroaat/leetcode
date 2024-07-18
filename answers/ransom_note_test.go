package answers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{ransomNote: "a", magazine: "b", expected: false},
		{ransomNote: "aa", magazine: "ab", expected: false},
		{ransomNote: "aa", magazine: "aab", expected: true},
	}

	for _, test := range tests {
		actual := canConstruct(test.ransomNote, test.magazine)
		assert.Equalf(t, test.expected, actual, "Expected %t, got %t", test.expected, actual)
	}
}
