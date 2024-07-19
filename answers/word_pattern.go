package answers

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)

	if len(words) != len(pattern) {
		return false
	}

	charToWord := make(map[byte]string)
	wordToChar := make(map[string]byte)

	for i := 0; i <= len(pattern)-1; i++ {
		char := pattern[i]
		word := words[i]

		if charToWord[char] != "" && charToWord[char] != word {
			return false
		}

		if wordToChar[word] != 0 && wordToChar[word] != char {
			return false
		}

		charToWord[char] = word
		wordToChar[word] = char
	}

	return true
}
