package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	conn, err1 := rpc.Dial("tcp", "127.0.0.1:8089")
	if err1 != nil {
		return
	}
	defer conn.Close()

	//接受返回值
	var reply string
	err2 := conn.Call("hello.SayHello", "wdf", &reply)
	if err2 != nil {
		return
	}

	fmt.Printf(reply)

}
