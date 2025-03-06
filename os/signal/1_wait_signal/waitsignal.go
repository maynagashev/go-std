package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(sigs chan os.Signal) {
	var (
		i        int
		shutdown bool
	)
	for !shutdown || i%5 != 0 {
		i++
		fmt.Printf("%d ", i)
		time.Sleep(1 * time.Second)
		select {
		case sig := <-sigs:
			fmt.Println("\r\nПолучен сигнал", sig)
			shutdown = true
		default:
			// select будет ожидать получения сигнала без default
		}
	}
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	worker(sigs)
	fmt.Println("Завершили работу")
}
