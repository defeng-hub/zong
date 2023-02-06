package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"

	"zong/mgo/day21/pb"
)

func main() {
	clientObj := &pb.String{Value: "1111111111啊1w1"}

	proto.Marshal(clientObj)
	// 序列化
	out, err := proto.Marshal(clientObj)
	if err != nil {
		log.Fatalln("Failed to encode obj:", err)
	}

	// 二进制编码
	fmt.Println("encode bytes: ", out)

	// 反序列化
	serverObj := &pb.String{}
	err = proto.Unmarshal(out, serverObj)
	if err != nil {
		log.Fatalln("Failed to decode obj:", err)
	}
	fmt.Println("decode obj: ", serverObj)
}

// encode bytes:  [10 12 104 101 108 108 111 32 112 114 111 116 111 51]
// decode obj:  value:"hello proto3"
