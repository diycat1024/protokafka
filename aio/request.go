package aio

import (
	"encoding/binary"
	comm "m1/common"
	pb "m1/protos"
)

type InReqMsg struct {
	Client *Client
	Cmd    pb.Cmd
	Data   []byte
}

func DecodeInReqMsg(b []byte) (*InReqMsg, error) {
	req := &InReqMsg{
		Cmd: pb.Cmd(binary.BigEndian.Uint32(b[4:comm.HeadLen])),
	}
	length := binary.BigEndian.Uint32(b[0:4])
	req.Data = b[comm.HeadLen : int(length)+comm.HeadLen]
	return req, nil
}
