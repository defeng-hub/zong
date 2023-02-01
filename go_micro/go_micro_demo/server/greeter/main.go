package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"greeter/handler"
	pb "greeter/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "greeter"
	version = "latest"
)

func main() {
	// 集成consul
	consulReg := consul.NewRegistry()
	
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Address("0.0.0.0:8080"), // 可以指定微服务ip,也可以不指定
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
	)

	// Register handler
	if err := pb.RegisterGreeterHandler(srv.Server(), new(handler.Greeter)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
