package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Hello struct {
}

// SayHello 第二个必须是指针类型
func (h Hello) SayHello(req string, res *string) error {
	*res = "你好," + req
	return nil
}
func main() {
	err1 := rpc.RegisterName("hello", new(Hello))
	if err1 != nil {
		fmt.Println(err1)
	}
	Listener, err2 := net.Listen("tcp", "127.0.0.1:8089")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer Listener.Close()

	for {
		fmt.Println("开始建立连接")
		conn, err3 := Listener.Accept()
		if err3 != nil {
			fmt.Println(err3)
		}
		rpc.ServeConn(conn)
	}
}
