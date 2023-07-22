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

type Stu struct {
	name   string
	yuwen  float64
	shuxue float64
	yingyu float64
	avg    float64
}
type cclass struct {
	students []*Stu
	yuwen    float64
	shuxue   float64
	yingyu   float64
}

func main() {
	// 结构体每个成员 是有默认值的
	user := User{time: time.Now()}
	fmt.Printf("[%s]-[%d]-[%f]-[%d]", user.name, user.age, user.height, user.time)

	s1 := Stu{
		name:   "s1",
		yuwen:  36,
		shuxue: 47,
		yingyu: 58,
	}
	s2 := Stu{
		name:   "s2",
		yuwen:  36,
		shuxue: 47,
		yingyu: 58,
	}
	s3 := Stu{
		name:   "s3",
		yuwen:  36,
		shuxue: 47,
		yingyu: 58,
	}
	cla := cclass{
		students: []*Stu{&s1, &s2, &s3},
	}

	fmt.Printf("%v", cla)

}
