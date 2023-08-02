package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	once := sync.Once{}
	// 实现一个单例模式
	once.Do(func() {
		m = sync.Map{}
	})
	m.Store("aaa", "111")
	v, _ := m.Load("aaa")
	fmt.Printf("%s", v)
}
