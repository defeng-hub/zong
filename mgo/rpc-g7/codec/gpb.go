package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

//obj2byte
func GobEncode(val interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	// 编码后放到encoder
	encoder := gob.NewEncoder(buf)
	//编码
	encoder.Encode(val)

	return buf.Bytes(), nil

}

//byte2obj
func GobDecode(data []byte, val interface{}) error {
	decoder := gob.NewDecoder(bytes.NewReader(data))

	return decoder.Decode(val)

}

type TestStruct struct {
	F1 string
	F2 int
}

func main() {
	bys, _ := GobEncode(&TestStruct{
		F1: "test",
		F2: 10,
	})
	fmt.Printf("%#v", bys)

	obj := TestStruct{}
	err := GobDecode(bys, &obj)
	if err != nil {
		return
	}
	fmt.Printf("\nxxx:%#v", obj)
}
