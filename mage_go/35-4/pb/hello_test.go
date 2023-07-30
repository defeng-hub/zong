package pbaaa_test

import (
	"fmt"
	"testing"
	pbaaa "zong/mage_go/35-4/pb"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestMarshal(t *testing.T) {
	p := &pbaaa.String{Value: "wdf"}
	should := assert.New(t)

	marshal, err := proto.Marshal(p)
	if should.NoError(err) {
		fmt.Println("marshal", marshal)
	}

	obj := pbaaa.String{}
	err = proto.Unmarshal(marshal, &obj)
	if should.NoError(err) {
		fmt.Println("obj:", obj.Value)
	}

}
