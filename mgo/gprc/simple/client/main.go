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
	fmt.Printf("%#v", resp)

}
