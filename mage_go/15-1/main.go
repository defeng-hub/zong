package main

import "fmt"

func slice_arg_1(arr []int) {
	arr[0] = 12

}

// 不定长参数
func func1(a int, b ...int) {
	fmt.Printf("%v\n", b)
}

func main() {
	arr1 := []int{1, 2, 3}
	slice_arg_1(arr1)
	fmt.Println(arr1)

	func1(1, 2, 3, 4, 5)

	//	append()就是不定长参数
	slice := append([]byte("abc"), "aaa"...)
	fmt.Println("slice", slice)

	//[]rune("王德丰")... 切片后边加三个点，就是把切片展开
	slice2 := append([]rune("abc"), []rune("王德丰")...)
	fmt.Println("slice2", slice2)

	ls := rune('王')
	slice3 := append([]rune("abc"), ls)
	fmt.Println("slice3", slice3)

	//	读map，
	//判断是否存在
	var mp = make(map[int]int, 20)
	mp[3] = 50
	if value, ok := mp[3]; ok {
		fmt.Printf("存在：%v", value)
	} else {
		fmt.Printf("不存在：map[%v]", 3)
	}
}
