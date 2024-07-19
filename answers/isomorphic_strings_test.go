package answers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{s: "egg", t: "add", expected: true},
		{s: "foo", t: "bar", expected: false},
		{s: "paper", t: "title", expected: true},
	}

	for _, test := range tests {
		actual := isIsomorphic(test.s, test.t)
		assert.Equalf(t, test.expected, actual, "Expected %t, got  %t.", test.expected, actual)
	}
}
