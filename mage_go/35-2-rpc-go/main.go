package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (h *HelloService) Hello(req string, res *string) error {
	*res = "hello" + req
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		return
	}
	for {
		conn, _ := listen.Accept()

		go rpc.ServeConn(conn)
	}
}
