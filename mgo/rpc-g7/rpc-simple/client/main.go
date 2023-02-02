package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	cilent, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		return
	}
	var reply string
	cilent.Call("HelloService.Hello", "wdf", &reply)
	fmt.Printf("xxx:%v\n", reply)
}
