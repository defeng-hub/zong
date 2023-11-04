package main

func isHappy(n int) bool {
	var m = make(map[int]int, 10)

	for {

		if v, ok := m[n]; ok {
			if v == 1 {
				return false
			}
		} else {
			m[n] = 1

		}

		zong := 0
		for n > 0 {
			y := n % 10
			zong += y * y
			n = n / 10
		}

		if zong == 1 {
			return true
		} else {
			n = zong
		}
	}
	return false
}
