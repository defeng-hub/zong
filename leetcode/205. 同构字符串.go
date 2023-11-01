package main

func isIsomorphic(s string, t string) bool {
	var smap = make(map[uint8]uint8, 10) //key s  value t
	var tmap = make(map[uint8]uint8, 10)
	for i := 0; i < len(s); i++ {
		v, ok := smap[s[i]]
		if ok {
			if v != t[i] {
				return false
			}
		} else {
			smap[s[i]] = t[i]
		}

	}
	for k, v := range smap {
		_, ok := tmap[v]
		if ok {
			return false
		} else {
			tmap[v] = k
		}

	}
	return true
}
