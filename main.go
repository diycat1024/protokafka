package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	pb "m1/protos"
	"net"
	"os"
	"os/signal"
	"syscall"
)
type server struct {
}

func doServerStuf(conn net.Conn)  {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if  err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("err reading: ", err.Error())
			return
		}
		if len <= 512 {
			fmt.Printf("Receive data: %v, len: %d\n", string(buf[:len]), len)
			msg := &pb.Msg{
				Head: &pb.Head{},
				Body: &pb.Body{},
			}
			err = proto.Unmarshal(buf[:len],msg)
			if err != nil {
					fmt.Printf("parse error : %v\n", err.Error())
					return
			}
			fmt.Printf("recv client : %s\n", msg.Body.GetData())
		}
	}
}

func main() {
	errs := make(chan error)

	l, err := net.Listen("tcp",":9091")
	if err != nil {
		errs <- err
	}

	go func() {
		os := make(chan os.Signal)
		signal.Notify(os, syscall.SIGINT)
		errs <- fmt.Errorf("%s\n", <-os)
	}()
	go func() {
		for  {
			conn, err := l.Accept()
			if err != nil {
				errs <- err
				break
			}
			go doServerStuf(conn)
		}
	}()


	fmt.Errorf("%s\n", <- errs)
}
