package main

import (
	"fmt"
	"m1/aio"
	"m1/logger"
	_ "m1/logic"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	defer logger.Logger.Sync()
	s := aio.NewServer()
	go func() {
		os := make(chan os.Signal)
		signal.Notify(os, syscall.SIGINT)
		s.Errs <- fmt.Errorf("%s\n", <-os)
	}()

	go s.Start()
	logger.Logger.Info("监听端口9091...")
	logger.Logger.Errorf("%s\n", <-s.Errs)
}
