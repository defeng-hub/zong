package main

import "fmt"

func main() {
	var m = make(map[string]int, 100)

	fmt.Println(m)
	m["aa"] = 3
	m["aa"] = 1
	m["bb"] = 2
	//m["cc"] = 2
	m["dd"] = 3
	m["ee"] = 4

	for key, value := range m {
		fmt.Printf("key:%s, value:%d\n", key, value)
	}

	delete(m, "bb")
	if value, ok := m["cc"]; ok {
		fmt.Printf("key: cc, value:%d", value)
	} else {
		fmt.Printf("不存在key:cc")
	}
}
