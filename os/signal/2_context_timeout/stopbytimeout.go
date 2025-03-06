package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	go func() {
		var i int
		for {
			fmt.Printf("%d ", i)
			i++
			time.Sleep(1 * time.Second)
		}
	}()
	<-ctx.Done()
	if ctx.Err() != nil {
		fmt.Printf("Ошибка:%v\n", ctx.Err())
	}
}
