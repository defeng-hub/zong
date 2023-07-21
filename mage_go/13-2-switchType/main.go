package main

import "fmt"

func main() {
	var num interface{} = "aa"

	// switch type
	//switch value := num.(type){}
	switch value := num.(type) {
	case int:
		fmt.Printf("int:%d", value)
	case string:
		fmt.Printf("string: %s", value)
	case float64:
		fmt.Printf("float: %.2f", value)
	}

	//	switch type 和 强制类型转换
	// num.(int)
	switch num.(type) {
	case int:
		value := num.(int)
		fmt.Printf("强制类型转换，int:%d", value)
	}
}
