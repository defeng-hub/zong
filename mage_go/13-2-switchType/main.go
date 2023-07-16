package main

import "fmt"

func main() {
	var num interface{} = "aa"
	switch value := num.(type) {
	case int:
		fmt.Println(value)
	case string:
		fmt.Printf("string: %s", value)
	}

}
