package main

import (
	"sync"
	"time"
)

type MyMap struct {
	mp         map[int]int
	keyToCh    map[int]chan struct{}
	sync.Mutex //是个结构体，不用人工初始化
}

func NewMyMap() *MyMap {
	return &MyMap{
		mp:      make(map[int]int),
		keyToCh: make(map[int]chan struct{}),
	}
}

func (m *MyMap) Get(key int, maxWaitingDuration time.Duration) (int, error) {
	m.Lock()
	value, ok := m.mp[key]
	if ok {
		m.Unlock()
		return value, nil
	}
	return 0, nil
}

func (m *MyMap) Put(key, value int) {
	m.Lock()
	defer m.Unlock()
	m.mp[key] = value
}

func main() {
	_ = make(chan int, 1)

}
