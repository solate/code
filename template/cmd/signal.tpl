package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

//根据信号停止程序
func SystemSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-ch:
			log.Print("stopping...")
			return
		default:
		}
	}
}
