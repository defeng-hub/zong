package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"greeter/proto/greeter"
	"net"
)

var _ greeter.GreeterServer = (*Hello)(nil)

type Hello struct {
}

func (h Hello) SayHello(c context.Context, req *greeter.HelloReq) (*greeter.HelloRes, error) {
	fmt.Println(req)
	return &greeter.HelloRes{Message: "你好" + req.Name}, nil
}

func main() {

	fmt.Println("aaa")
	grpcServer := grpc.NewServer()
	greeter.RegisterGreeterServer(grpcServer, &Hello{})
	listen, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		return
	}
	defer listen.Close()
	grpcServer.Serve(listen)
}
