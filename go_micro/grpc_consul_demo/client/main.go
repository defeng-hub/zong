package main

import (
	"client/proto/greeter"
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // 非常重要,用来负载均衡,直接使用consul://
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {
	conn, _ := grpc.Dial("consul://192.168.88.135:8500/HelloService",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 负载均衡,轮询的方式
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)

	//// 原始方法,注册consul服务
	//consulConfig := api.DefaultConfig()
	////客户机的地址, 也可以是consul的服务端,但是一般用客户端
	//consulConfig.Address = "192.168.88.135:8500"
	//
	//consulClient, _ := api.NewClient(consulConfig)
	//serviceEntry, _, _ := consulClient.Health().Service("HelloService", "", false, nil)
	//address := serviceEntry[0].Service.Address + ":" + strconv.Itoa(serviceEntry[0].Service.Port)
	//
	//fmt.Println("address", address)
	//// 连接服务器
	//conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	return
	//}
	client := greeter.NewGreeterClient(conn)

	for {
		var res *greeter.HelloRes

		res, err := client.SayHello(context.Background(), &greeter.HelloReq{Name: "wdf"})
		if err != nil {
			return
		}
		fmt.Printf("%#v\n", res.Message)
		time.Sleep(time.Millisecond)
	}

}
