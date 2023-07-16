package main

import (
	"fmt"
)

func main() {
	func1()
}

func func1() {
	var i int64 = 90
	var b = byte(i)

	i = int64(b)
	fmt.Println("i=", i, ";b=", b)
}
