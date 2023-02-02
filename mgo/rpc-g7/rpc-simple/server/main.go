package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (h *HelloService) Hello(req string, reply *string) error {
	*reply = "hello:" + req
	return nil
}
func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	listen, _ := net.Listen("tcp", ":1234")
	for {
		conn, _ := listen.Accept()
		go rpc.ServeConn(conn)
	}
}
