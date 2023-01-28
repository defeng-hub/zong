package main

import (
	"context"
	"fmt"
	api "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"greeter/proto/greeter"
	"net"
)

var _ greeter.GreeterServer = (*Hello)(nil)

type Hello struct {
}

func (h Hello) SayHello(c context.Context, req *greeter.HelloReq) (*greeter.HelloRes, error) {
	fmt.Println(req)
	return &greeter.HelloRes{Message: "你好," + req.Name}, nil
}

func main() {
	//注册consul服务
	consulConfig := api.DefaultConfig()
	//客户机的地址, 也可以是consul的服务端,但是一般用客户端
	consulConfig.Address = "192.168.88.135:8500"

	consulClient, _ := api.NewClient(consulConfig)

	// 配置注册服务的参数
	agentService := api.AgentServiceRegistration{
		ID:      "1",
		Tags:    []string{},
		Name:    "HelloService",
		Port:    8081,
		Address: "192.168.1.25",
		Check: &api.AgentServiceCheck{
			TCP:      "192.168.1.25:8081",
			Timeout:  "60s",
			Interval: "5s",
		},
	}
	err := consulClient.Agent().ServiceRegister(&agentService)
	if err != nil {
		return
	}

	grpcServer := grpc.NewServer()
	greeter.RegisterGreeterServer(grpcServer, &Hello{})
	listen, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		return
	}
	defer listen.Close()
	grpcServer.Serve(listen)
}
