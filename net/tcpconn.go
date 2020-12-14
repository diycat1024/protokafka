package net

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	comm "m1/common"
	pb "m1/protos"
	"net"
)

type ConnCtx struct {
	ClientConn net.Conn
	Scanner    *bufio.Scanner
	Server     *server
}

func NewConnCtx(c net.Conn) *ConnCtx {
	return &ConnCtx{
		ClientConn: c,
		Scanner:    bufio.NewScanner(c),
	}
}

func (c *ConnCtx) ParseData() {
	defer c.ClientConn.Close()
	for {
		c.Scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if !atEOF {
				if len(data) > comm.HeadLen { //4字节数据包长度  4字节指令
					length := int32(0)
					binary.Read(bytes.NewReader(data[0:4]), binary.BigEndian, &length)
					if length <= 0 {
						return 0, nil, fmt.Errorf("length is 0")
					}
					fmt.Printf("len_data %d; length: %d\n", len(data), length)
					if int(length)+comm.HeadLen <= len(data) {
						return int(length) + comm.HeadLen , data[: int(length)+ comm.HeadLen], nil
					}
				}
			}
			return
		})

		for c.Scanner.Scan() {
			fmt.Println("scanner msg: ", string(c.Scanner.Bytes()))
			length := binary.BigEndian.Uint32(c.Scanner.Bytes()[0:4])
			cmd := binary.BigEndian.Uint32(c.Scanner.Bytes()[4:comm.HeadLen])
			switch pb.MsgID(cmd) {
			case pb.MsgID_eMsgToLSFromGC_AskLogin:
				msg := &pb.AskLogin{}
				err := proto.Unmarshal(c.Scanner.Bytes()[comm.HeadLen:int(length) + comm.HeadLen], msg)
				if err != nil {
					fmt.Printf("proto Unmarshal err: %s\n", err.Error())
					return
				}
				fmt.Printf("msg.msgid: %d, msg.session: %s", msg.Msgid, msg.Sessionid)

				//todo 登录成功


			default:
				fmt.Println("unknow  !!!")
			}

		}

		if err := c.Scanner.Err(); err != nil {
			fmt.Println("无数据包！")
			break
		}
	}
}
