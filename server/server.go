package server

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

type server struct {
	Errs chan error
}

func NewServer() *server {
	return &server{
		Errs: make(chan error),
	}
}

func (s *server)doServerStuf(conn net.Conn)  {
	defer conn.Close()
	for {
		scanner := bufio.NewScanner(conn)
		scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if !atEOF {
				if len(data) > comm.HeadLen  {  //4字节数据包长度  4字节指令
					length := int32(0)
					binary.Read(bytes.NewReader(data[0:4]),binary.BigEndian, &length)
					if length <= 0 {
						return 0, nil, fmt.Errorf("length is 0")
					}
					fmt.Printf("len_data %d; length: %d\n", len(data), length)
					if int(length) + comm.HeadLen <= len(data) {
						return int(length), data[comm.HeadLen:int(length) + comm.HeadLen], nil
					}
				}
			}
			return
		})

		for scanner.Scan() {
			fmt.Println("scanner msg: ", string(scanner.Bytes()))
			msg := &pb.AskLogin{}
			s := bytes.Buffer{}
			s.Write(scanner.Bytes())
			s.Write([]byte("  "))
			err := proto.Unmarshal(s.Bytes(), msg)
			if err != nil {
				fmt.Printf("safasdfcccccccc: %s\n", err.Error())
				return
			}
			fmt.Println("msg.msgid=", msg.Msgid)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("无数据包！")
		}

	}
}

func (s *server)Accept() {

	l, err := net.Listen("tcp",":9091")
	if err != nil {
		s.Errs <- err
	}

	go func() {
		conn ,err := l.Accept()
		if err != nil {
			s.Errs <- err
		}
		go s.doServerStuf(conn)
	}()
}
