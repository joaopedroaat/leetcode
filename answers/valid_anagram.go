package answers

import "fmt"

func IsAnagramTestCase() {
	res := IsAnagram("anagram", "nagaram")
	fmt.Println(res)
}

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make(map[byte]int)
	m2 := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		m1[s[i]]++
		m2[t[i]]++
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}

	return true
}
