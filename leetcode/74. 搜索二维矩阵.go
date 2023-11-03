package main

func searchMatrix(matrix [][]int, target int) bool {
	var idx = 0
	for i, ints := range matrix {
		if ints[len(ints)-1] < target {
			continue
		}
		idx = i
		break
	}
	for _, v := range matrix[idx] {
		if v == target {
			return true
		}
	}
	return false
}
