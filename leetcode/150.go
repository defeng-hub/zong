package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	}

	var p1 = 0
	var p2 = 1

	for p2 < n {
		if nums[p1] != nums[p2] {
			nums[p1] = nums[p2]
			p1++
		}
	}

	return p1 + 1
}
func func150() {

}
