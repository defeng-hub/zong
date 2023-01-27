package main

import (
	"fmt"
	"net/rpc"
)

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

func main() {
	conn, err1 := rpc.Dial("tcp", "127.0.0.1:8020")
	if err1 != nil {
		return
	}
	defer conn.Close()

	//接受返回值
	var reply AddGoodsRes
	err2 := conn.Call("goods.AddGoods", AddGoodsReq{
		Id:      2,
		Title:   "拖鞋666",
		Price:   100,
		Content: "666",
	}, &reply)
	if err2 != nil {
		return
	}
	fmt.Println(reply)

	var reply2 GetGoodsRes
	err3 := conn.Call("goods.GetGood", GetGoodsReq{
		Id: 3,
	}, &reply2)
	if err3 != nil {
		return
	}
	fmt.Println(reply2)

}
