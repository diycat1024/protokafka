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

type Client struct {
	ClientConn net.Conn
	Scanner    *bufio.Scanner
}

func NewClient(c net.Conn) *Client {
	return &Client{
		ClientConn: c,
		Scanner:    bufio.NewScanner(c),
	}
}

func (c *Client) doServerStuf() {
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
						return int(length), data[comm.HeadLen : int(length)+comm.HeadLen], nil
					}
				}
			}
			return
		})

		for c.Scanner.Scan() {
			fmt.Println("scanner msg: ", string(c.Scanner.Bytes()))
			msg := &pb.AskLogin{}
			s := bytes.Buffer{}
			s.Write(c.Scanner.Bytes())
			s.Write([]byte("  "))
			err := proto.Unmarshal(s.Bytes(), msg)
			if err != nil {
				fmt.Printf("proto Unmarshal err: %s\n", err.Error())
				return
			}
			fmt.Printf("msg.msgid: %d, msg.session: %s", msg.Msgid, msg.Sessionid)
		}

		if err := c.Scanner.Err(); err != nil {
			fmt.Println("无数据包！")
			break
		}
	}
}
