package main

import "fmt"

func f1(i interface{}) {
	if v, ok := i.(int); ok {
		fmt.Println("整数", v)
	}
}

func main() {
	color := "yellow"

	switch color {
	case "red":
		fmt.Println("红")
	case "yellow":
		fmt.Println("黄")
		fallthrough
	case "green":
		fmt.Println("绿")
	default:
		fmt.Println("other")
	}
}
