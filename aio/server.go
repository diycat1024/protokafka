package aio

import (
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"m1/logger"
	pb "m1/protos"
	"net"
	"sync"
)

type LogicHandler interface {
	HandlerCommon(msg *InReqMsg)
	HandlerError(cmd pb.Cmd, msg *proto.Message)
	HandlerResponse(cmd pb.Cmd, msg *proto.Message)
}

var LoginHandlerMap sync.Map

type server struct {
	Errs    chan error
	Clients *sync.Map

	Listener   net.Listener
	ClientJoin chan net.Conn
	ClientQuit chan *Client
	InputMsg   chan InReqMsg
}

func NewServer() *server {
	s := &server{
		Errs:       make(chan error),
		Clients:    &sync.Map{},
		ClientJoin: make(chan net.Conn),
		ClientQuit: make(chan *Client),
		InputMsg:   make(chan InReqMsg),
	}
	go s.listen()
	return s
}

func (s *server) listen() {
	for {
		select {
		case msg := <-s.InputMsg:
			s.RecvHandler(msg)
		case conn := <-s.ClientJoin:
			s.JoinHandler(conn)
		case client := <-s.ClientQuit:
			s.QuitHandler(client)
		}

	}
}

func (s *server) Start() {
	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		s.Errs <- err
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			s.Errs <- err
		}
		s.ClientJoin <- conn
	}
}

func (s *server) AddClient(client *Client) {
	s.Clients.Store(client.Key, client)
}

func (s *server) GetClient(key string) *Client {
	if client, b := s.Clients.Load(key); b {
		return client.(*Client)
	}
	return nil
}

func (s *server) DelClient(key string) {
	s.Clients.Delete(key)
}

func (s *server) RecvHandler(msg InReqMsg) {
	logger.Logger.Debug("a msg recv")
	h, b := LoginHandlerMap.Load(msg.Cmd)
	if !b {
		logger.Logger.Debugf("msg.cmd not register %d \n", msg.Cmd)
		return
	}
	//todo have bug
	handler := *h.(*LogicHandler)
	handler.HandlerCommon(&msg)
}

func (s *server) JoinHandler(conn net.Conn) {
	logger.Logger.Info("a conn add")
	key := uuid.New()
	client := NewClient(key, conn)
	s.AddClient(client)
	go func() {
		for {
			//循环读取消息
			msg := <-client.Input
			s.InputMsg <- msg
		}
	}()
	// 开启协程 等待断开
	go func() {
		conn := <-client.Quit
		s.ClientQuit <- conn
	}()
	msg := &pb.AskLogin{}
	msg.Cmd = pb.Cmd_eMsgToLSFromGC_AskLogin
}

func (s *server) QuitHandler(conn *Client) {
	logger.Logger.Debug("a client quit")
	s.Clients.Delete(conn.Key)
}
