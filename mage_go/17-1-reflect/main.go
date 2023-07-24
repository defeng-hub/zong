package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
}

func func12() {
	typeI := reflect.TypeOf('a')
	typeS := reflect.TypeOf("hello")
	fmt.Println(typeI.Kind() == reflect.Chan)
	fmt.Println(typeS.String())
	fmt.Println(typeS.String())
	fmt.Println(typeS.String())

}

func func2() {
	u := User{}
	ref := reflect.TypeOf(u)
	fmt.Println(ref.String())
	fmt.Println(ref.Kind())

	//传入指针
	u2 := &User{}
	ref2 := reflect.TypeOf(u2)
	fmt.Println(ref2.String())
	fmt.Println(ref2.Kind())

	// elem 从指针转为元素
	ref3 := ref2.Elem()
	//ref3 := reflect.TypeOf(u3)
	fmt.Println(ref3.String())
	fmt.Println(ref3.Kind())

}

func main() {
	//func1()
	func2()
}
