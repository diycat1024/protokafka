package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "m1/protos"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		fmt.Errorf("%s\n", err.Error())
		return
	}
	msg := &pb.AskLogin{
		Cmd:  pb.Cmd_eMsgToLSFromGC_AskLogin,
		Plat: "1",
		Uid:  "1001",
	}

	msgbyte, err := proto.Marshal(msg)
	if err != nil {
		fmt.Errorf("msg Marshal error %s\n", err.Error())
		return
	}

	buf := &bytes.Buffer{}
	var head []byte
	head = make([]byte, 8)
	binary.BigEndian.PutUint32(head[0:4], uint32(bytes.Count(msgbyte, nil)-1))
	binary.BigEndian.PutUint32(head[4:8], uint32(pb.Cmd_eMsgToLSFromGC_AskLogin))
	buf.Write(head[:8])
	buf.Write(msgbyte)
	fmt.Printf("%v\n", string(buf.Bytes()))

	conn.Write(buf.Bytes())

	defer conn.Close()
}
