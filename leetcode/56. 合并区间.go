package main

import (
	"sort"
)

func merge1(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int
	result = append(result, intervals[0])

	for _, e := range intervals[1:] {
		if result[len(result)-1][1] < e[0] {
			result = append(result, e)
		} else {
			result[len(result)-1][1] = max(e[1], result[len(result)-1][1])
		}
	}
	return nil
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
