package answers

func isIsomorphic(s string, t string) bool {
	m1, m2 := make(map[byte]byte), make(map[byte]byte)

	for i := 0; i <= len(s)-1; i++ {
		a, b := s[i], t[i]
		if m1[a] == 0 && m2[b] == 0 {
			m1[a] = b
			m2[b] = a
		} else if m1[a] != b || m2[b] != a {
			return false
		}
	}

	return true
}
