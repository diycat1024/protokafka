package net

import (
	"net"
)

type server struct {
	Errs    chan error
	Clients map[string]*ConnCtx
	Login   chan ConnCtx
}

func NewServer() *server {
	return &server{
		Errs:    make(chan error),
		Clients: make(map[string]*ConnCtx),
		Login: make(chan ConnCtx, 1),
	}
}

func (s *server) Accept() {

	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		s.Errs <- err
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			s.Errs <- err
		}
		connctx := NewConnCtx(conn)
		go connctx.ParseData()
	}
}
