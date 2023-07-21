package main

import (
	"fmt"
	"time"
)

type User struct {
	name   string
	age    int
	height float64
	time   time.Time
}
type usermap map[int]string

//可以跟任何类型添加方法，比如给map加方法

func (um usermap) getaa() {

}
func main() {
	// 结构体每个成员 是有默认值的
	user := User{time: time.Now()}
	fmt.Printf("[%s]-[%d]-[%f]-[%d]", user.name, user.age, user.height, user.time)

}
