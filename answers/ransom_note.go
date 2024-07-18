package answers

func canConstruct(ransomNote string, magazine string) bool {
	// Magazine chars map
	m1 := make(map[rune]int)
	// Ransom note chars map
	m2 := make(map[rune]int)

	// Count chars on magazine
	for _, rune := range magazine {
		m1[rune] = m1[rune] + 1
	}

	// Count chars on ransom note
	for _, rune := range ransomNote {
		m2[rune] = m2[rune] + 1
	}

	// Compare chars counts
	for key, value := range m2 {
		if m1[key] < value {
			return false
		}
	}

	return true
}
