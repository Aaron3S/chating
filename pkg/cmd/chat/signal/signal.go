package signal

import (
	"os"
	"os/signal"
	"syscall"
)

var once = make(chan struct{})

func SetUpSignalHandler() (stopCh <-chan struct{}) {
	close(once)
	stop := make(chan struct{})
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1)
	}()
	return stop
}
