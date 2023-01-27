package main

import (
	"client/proto/greeter"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("client")

	// 连接服务器
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	client := greeter.NewGreeterClient(conn)

	var res *greeter.HelloRes
	res, err = client.SayHello(context.Background(), &greeter.HelloReq{Name: "wdf"})
	if err != nil {
		return
	}
	fmt.Printf("%#v\n", res.Message)
}
