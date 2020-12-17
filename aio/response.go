package aio

import (
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	pb "m1/protos"
)

type OutResMsg struct {
	//Client *Client        `json:"_"`
	Cmd  pb.Cmd         `json:"cmd"`
	Data *proto.Message `json:"data"`
}

func NewOutResMsg(cmd pb.Cmd, data proto.Message) *OutResMsg {
	return &OutResMsg{
		Cmd:  cmd,
		Data: &data,
	}
}

func (resp *OutResMsg) Decode() []byte {
	msgbyte, err := proto.Marshal(*resp.Data)
	if err != nil {
		return nil
	}
	buf := &bytes.Buffer{}
	var head []byte
	head = make([]byte, 8)
	binary.BigEndian.PutUint32(head[0:4], uint32(bytes.Count(msgbyte, nil)-1))
	binary.BigEndian.PutUint32(head[4:8], uint32(resp.Cmd))
	buf.Write(head[:8])
	buf.Write(msgbyte)
	return buf.Bytes()
}
