// main包
package main

import "fmt"

/*
  var a int=5
  var a int=5
  var a int=5
	var a int=5

*/
var B int = 3

func bitDemo() {
	//NOTE:函数没有返回值
	var str string = "a啊hello" //TODO:线程不安全
	for _, ele := range str {
		fmt.Printf("%d,%c,%T\n", ele, ele, ele)
	}
}
func main() {
	bitDemo()
}
