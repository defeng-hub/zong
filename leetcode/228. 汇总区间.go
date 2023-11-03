package main

import "fmt"

func summaryRanges(nums []int) []string {
	var resp []string

	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{fmt.Sprintf("%d", nums[1])}
	}

	slow := 0
	fast := 1

	for slow < len(nums) {
		if nums[fast]-1 == nums[fast-1] {
			fast++
		} else {
			if fast-1 > slow {
				resp = append(resp, fmt.Sprintf("%d->%d", nums[slow], nums[fast-1]))
				slow = fast
			} else {
				resp = append(resp, fmt.Sprintf("%d", nums[slow]))
				slow++
			}
			fast++
		}
	}
	return resp
}
