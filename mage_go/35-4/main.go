package main

import (
	"sync"
)

type singleton struct {
}

var instance *singleton
var once sync.Once

func main() {
	instance = new(singleton)
	println(instance)
	instance = new(singleton)
	println(instance)
	return
}
