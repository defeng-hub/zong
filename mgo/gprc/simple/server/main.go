package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gprc/simple/server/pb"
	"log"
	"net"
)

type HelloServiceServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloServiceServer) Hello(ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	return &pb.Response{Value: fmt.Sprintf("helloo, %s", req.Value)}, nil
}
func (s *HelloServiceServer) Channel(req pb.HelloService_ChannelServer) error {
	for {
		args, _ := req.Recv()
		req.Send(&pb.Response{Value: fmt.Sprintf("helloo,%s", args.GetValue())})
		//return nil
	}
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, new(HelloServiceServer))

	lis, _ := net.Listen("tcp", "127.0.0.1:1234")

	log.Printf("listen 启动成功")
	grpcServer.Serve(lis)
}
