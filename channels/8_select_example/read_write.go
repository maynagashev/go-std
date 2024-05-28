package main

import "fmt"

func main() {
	chIn := make(chan int)
	chOut := make(chan int)
	quit := make(chan struct{})

	go func() {
		for i := 0; i < 15; i++ {
			chIn <- i
		}
		close(chIn)
	}()
	go func() {
		for x := range chIn {
			chOut <- x * 2
		}
		close(chOut)
	}()
	go func() {
		for x := range chOut {
			fmt.Printf("%d ", x)
		}
		quit <- struct{}{}
	}()
	<-quit
}
