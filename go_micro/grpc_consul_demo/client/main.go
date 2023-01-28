package main

import (
	"client/proto/greeter"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

func main() {
	//注册consul服务
	consulConfig := api.DefaultConfig()
	//consulConfig.Address = "127.0.0.1:8500"

	consulClient, _ := api.NewClient(consulConfig)
	serviceEntry, _, _ := consulClient.Health().Service("HelloService", "", false, nil)
	address := serviceEntry[0].Service.Address + ":" + strconv.Itoa(serviceEntry[0].Service.Port)
	fmt.Println("address", address)
	// 连接服务器
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
