package main

import (
	"fmt"
	"m1/aio"
	"m1/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger.InitLogger()
	defer logger.Logger.Sync()
	s := aio.NewServer()

	go func() {
		os := make(chan os.Signal)
		signal.Notify(os, syscall.SIGINT)
		s.Errs <- fmt.Errorf("%s\n", <-os)
	}()
	fmt.Errorf("%s\n", <-s.Errs)
}
