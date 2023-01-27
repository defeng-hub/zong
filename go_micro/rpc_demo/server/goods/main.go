package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Goods struct {
}
type AddGoodsReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}
type AddGoodsRes struct {
	Success bool
	Message string
}
type GetGoodsReq struct {
	Id int
}
type GetGoodsRes struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

func (this Goods) AddGoods(req AddGoodsReq, res *AddGoodsRes) error {
	fmt.Println("添加了:", req)
	*res = AddGoodsRes{
		Success: true,
		Message: "添加成功",
	}
	return nil
}
func (this Goods) GetGood(req GetGoodsReq, res *GetGoodsRes) error {
	*res = GetGoodsRes{
		Id:      req.Id,
		Title:   "测试",
		Price:   70,
		Content: "1111xxx嘻嘻嘻",
	}
	return nil
}
func main() {
	fmt.Printf("goods server")
	err1 := rpc.RegisterName("goods", new(Goods))
	if err1 != nil {
		fmt.Println(err1)
	}
	listener, err2 := net.Listen("tcp", "127.0.0.1:8020")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer listener.Close()

	for {
		fmt.Println("开始建立连接")
		conn, err3 := listener.Accept()
		if err3 != nil {
			return
		}
		rpc.ServeConn(conn)
	}
}
