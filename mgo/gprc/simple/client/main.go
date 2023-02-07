package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"simple/server/pb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
		return
	}

	client := pb.NewHelloServiceClient(conn)
	resp, err := client.Hello(context.Background(), &pb.Request{Value: "alice"})
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("aaa:%#v\n", resp)

	stream, err := client.Channel(context.Background())
	if err != nil {
		panic(err)
	}

	go func() {
		stream.Send(&pb.Request{Value: "alice"})
		if err != nil {
			panic(err)
		}
	}()

	recv, err := stream.Recv()
	if err != nil {
		panic(err)
	}
	fmt.Printf("bbb:%#v\n", recv)

}
