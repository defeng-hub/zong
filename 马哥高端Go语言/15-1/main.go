package main

import "fmt"

func slice_arg_1(arr []int) {
	arr[0] = 12

}

func main() {
	arr1 := []int{1, 2, 3}
	slice_arg_1(arr1)

	fmt.Println(arr1)
}
