package main

func findPeakElement(nums []int) int {
	min := -2147483648
	nums = append([]int{min}, nums...)
	nums = append(nums, min)
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] <= i && nums[i+1] <= i {
			return i - 1
		}
	}

	return 0
}
