package main

import (
	"fmt"
	"m1/net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := net.NewServer()

	go func() {
		os := make(chan os.Signal)
		signal.Notify(os, syscall.SIGINT)
		s.Errs <- fmt.Errorf("%s\n", <-os)
	}()
	fmt.Errorf("%s\n", <-s.Errs)
}
