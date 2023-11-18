package main

func merge1(intervals [][]int) [][]int {
	var list = [10000]int{}

	for _, interval := range intervals {
		for i := interval[0]; i <= interval[1]; i++ {
			list[i] = 1
		}
	}

	var res [][]int
	start := 0
	end := 0
	for i := 1; i < len(list); i++ {
		if list[i] == 1 {
			start = i
			end = i
			for list[end] == 1 {
				end++
			}
			i = end
			end--
			res = append(res, []int{start, end})
		}
	}
	return res
}
