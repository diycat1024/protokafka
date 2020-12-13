package net

import (
	"net"
)

type server struct {
	Errs    chan error
	Clients map[string]*Client
}

func NewServer() *server {
	return &server{
		Errs:    make(chan error),
		Clients: make(map[string]*Client),
	}
}

func (s *server) Accept() {

	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		s.Errs <- err
	}

	go func() {
		conn, err := l.Accept()
		if err != nil {
			s.Errs <- err
		}
		client := NewClient(conn)
		go client.doServerStuf()
	}()
}
