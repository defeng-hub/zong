package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"proto_demo/proto/userService"
)

func main() {
	fmt.Println("aaa")

	userinfo := &userService.Userinfo{
		Username: "王德丰",
		Age:      18,
		Hobby:    []string{"电脑", "游戏", "学习"},
	}
	fmt.Println(userinfo)

	data, err := proto.Marshal(userinfo)
	if err != nil {
		return
	}
	fmt.Println("数据", data)

	user := &userService.Userinfo{}
	proto.Unmarshal(data, user)

	fmt.Println("user", user)
}
