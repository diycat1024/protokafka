package main

import (
	"bufio"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"m1/aio"
	"m1/logger"
	pb "m1/protos"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func RecvMsg(msg *aio.InReqMsg) {
	switch msg.Cmd {
	case pb.Cmd_eMsgToLSFromGC_SendGroupMsg:
		m := &pb.SendGroupMsg{}
		if err := proto.Unmarshal(msg.Data, m); err != nil {
			logger.Logger.Errorf("%s\n", err.Error())
			return
		}
		fmt.Printf("recv msg: %s\n", m.String())
	case pb.Cmd_eMsgToLSFromGC_AskLogin:
		m := &pb.AskLogin{}
		if err := proto.Unmarshal(msg.Data, m); err != nil {
			logger.Logger.Errorf("%s\n", err.Error())
			return
		}
		fmt.Printf("recv msg: %s\n", m.String())
	default:
		logger.Logger.Debugf("recv a unknown data!!!")

	}
}

func SendMsg(client *aio.Client) {
	msg := &pb.SendGroupMsg{
		Cmd:  pb.Cmd_eMsgToLSFromGC_SendGroupMsg,
		Send: "101",
		Recv: "102",
		Plat: "1",
		Data: "hi I'm 101!!!!!",
	}
	client.PutOut(aio.NewOutResMsg(msg.Cmd, msg))
}

func AskLogin(client *aio.Client) {
	msg := &pb.AskLogin{
		Cmd:  pb.Cmd_eMsgToLSFromGC_AskLogin,
		Plat: "1",
		Uid:  "1001",
	}

	client.PutOut(aio.NewOutResMsg(msg.Cmd, msg))
	defer client.Quiting()
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		fmt.Errorf("%s\n", err.Error())
		return
	}
	key := uuid.New()
	client := aio.NewClient(key, conn)
	go func() {
		for {
			//接收消息
			m := <-client.Input
			RecvMsg(m)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	logger.Logger.Info("Simple Shell")
	logger.Logger.Info("---------------------")

	go func() {
		for {
			logger.Logger.Debugf("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text = strings.Replace(text, "\n", "", -1)

			if strings.Compare("hi", text) == 0 {
				fmt.Println("hello, Yourself")
			}
			switch text {
			case "hello":
				fmt.Println("hello, Yourself")
			case "login":
				//登录
				AskLogin(client)
			case "send":
				//发送消息
				SendMsg(client)
			}

		}
	}()

	errs := make(chan error)
	go func() {
		o := make(chan os.Signal)
		signal.Notify(o, syscall.SIGINT)
		errs <- fmt.Errorf("%s\n", <-o)
	}()

	logger.Logger.Info("success conn to 127.0.0.1:9091")
	logger.Logger.Errorf("%s\n", <-errs)
}
