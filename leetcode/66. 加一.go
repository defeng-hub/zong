package main

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	// digits 中所有的元素均为 9
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}
