package main

import (
	"fmt"
	"strconv"
)

type user struct {
	name string
	age  int
}
type abc = user

func (self abc) String() string {
	return self.name + "---" + strconv.Itoa(self.age)
}
func (self user) String2() string {
	return self.name + "--2--" + strconv.Itoa(self.age)
}

func scope() {
	var b byte = 255
	//var a int8 = 127
	//var c uint = 256

	var f float32 = 2.4567
	var t bool
	var s string
	var r rune = 9787
	var arr []int

	var obj abc = abc{
		name: "wdf",
		age:  18,
	}

	var obj2 abc
	fmt.Printf("default value, byte: %d\n", b)
	fmt.Printf("default value, float32: %f,%.2f,%.3e,%g\n", f, f, f, f)
	fmt.Printf("default value, bool: %t\n", t)
	fmt.Printf("default value, string: [%s]\n", s)
	fmt.Printf("default value, rune: %d, [%c]\n", r, r)
	fmt.Printf("default value, slice: %v, [%t]\n", arr, arr == nil)
	fmt.Printf("default value, obj: %v, [%t]\n", obj, obj == abc{name: "wdf", age: 18})
	fmt.Printf("default value, obj: %v, [%t]\n", obj2.String2(), obj == obj2)

}
func main() {
	scope()
}
