package main

func isAnagram(s string, t string) bool {
	var m = make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		_, ok := m[s[i]]
		if !ok {
			m[s[i]] = 1
		} else {
			m[s[i]]++
		}
	}

	for i := 0; i < len(t); i++ {
		v, ok := m[t[i]]
		if !ok || v <= 0 {
			return false
		} else {
			m[t[i]]--
		}
	}
	for _, i := range m {
		if i != 0 {
			return false
		}
	}

	return true
}
