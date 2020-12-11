package main

import (
	"fmt"
	"m1/protos"
	"net"
)

const (
	login  = 0x001
	logout = 0x002
)

func main() {
	conn, err := net.Dial("tcp","127.0.0.1:9091")
	if err != nil {
		fmt.Errorf("%s\n", err.Error())
		return
	}

	head := &protos.Head{
		Version: 1,
		Cmd: login,
		Len: 0,
		Encry: 0,
		Ext1: 0,
		Ext2: 0,
	}
	body := &protos.Body{
		Bodylen: 10,
		Data: "hello world",
	}
	msg :=  protos.Msg{
		Head: head,
		Body: body,
	}
	msg.Head.Len = uint64(len([]byte(msg.String())))
	fmt.Printf("len: %d, msg: %s\n", msg.Head.Len, msg.String())
	conn.Write([]byte(msg.String()))
	defer conn.Close()
}
