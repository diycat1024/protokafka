package main

import (
	"fmt"
	"m1/server"
	"os"
	"os/signal"
	"syscall"
)


func main() {
	s := server.NewServer()
	go s.Accept()

	errs := make(chan error)

	go func() {
		os := make(chan os.Signal)
		signal.Notify(os, syscall.SIGINT)
		errs <- fmt.Errorf("%s\n", <-os)
	}()
	fmt.Errorf("%s\n", <- s.Errs)
}
