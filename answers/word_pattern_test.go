package answers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordPattern(t *testing.T) {
	tests := []struct {
		pattern  string
		s        string
		expected bool
	}{
		{pattern: "abba", s: "dog cat cat dog", expected: true},
		{pattern: "abba", s: "dog cat cat fish", expected: false},
		{pattern: "aaaa", s: "dog cat cat dog", expected: false},
	}

	for _, test := range tests {
		actual := wordPattern(test.pattern, test.s)
		assert.Equalf(t, test.expected, actual, "Expected %t, got %t.", test.expected, actual)
	}
}
