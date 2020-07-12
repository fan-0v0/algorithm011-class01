func isAnagram(s string, t string) bool {
	m := make(map[rune]int)
	for _, c := range s {
		_, ok := m[c]
		if ok {
			m[c]++
		} else {
			m[c] = 1
		}
	}
	for _, c := range t {
		_, ok := m[c]
		if ok {
			m[c]--
		} else {
			m[c] = -1
		}
	}
	for _, val := range m {
		if val != 0 {
			return false
		}
	}
	return true
}