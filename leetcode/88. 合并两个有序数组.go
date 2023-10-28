package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}

		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1[0:], sorted)
}

func fun88() {
	var nums1 []int
	var nums2 []int

	nums1 = []int{1, 2, 3, 4, 5}
	nums2 = []int{6, 7, 8, 9, 0}
	//fmt.Println("nums1[3:]", nums1[3:])
	//
	//i := copy(nums1[3:], nums2)
	//fmt.Println(i)
	//fmt.Println(nums1)

	merge(nums1, 3, nums2, 3)

}
