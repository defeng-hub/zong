package main

func majorityElement(nums []int) int {
	n := len(nums)
	m := make(map[int]int, 10)

	for _, v := range nums {
		_, ok := m[v]
		if ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for k, v := range m {
		if v > n/2 {
			return k
		}
	}
	return 0
}
func func169() {

}
