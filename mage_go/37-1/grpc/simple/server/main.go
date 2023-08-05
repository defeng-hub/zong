package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	pbaaa "simple/server/pb"
)

type HelloServiceServer struct {
	pbaaa.UnimplementedHelloServiceServer
}

func (HelloServiceServer) Hello(ctx context.Context, req *pbaaa.Request) (*pbaaa.Response, error) {
	return &pbaaa.Response{Value: "666777;" + req.Value}, nil
}

func main() {
	server := grpc.NewServer()
	pbaaa.RegisterHelloServiceServer(server, new(HelloServiceServer))

	listen, _ := net.Listen("tcp", ":1234")
	server.Serve(listen)
}
