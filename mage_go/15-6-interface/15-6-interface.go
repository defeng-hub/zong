package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	once := sync.Once{}
	once.Do(func() {
		m = sync.Map{}
	})
	m.Store("aaa", "111")
	v, _ := m.Load("aaa")
	fmt.Printf("%s", v)
}
