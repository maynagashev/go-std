package main

import "fmt"

func main() {
	//var ch chan int // nil channel locks on write
	ch := make(chan int)

	go func() {
		ch <- 7
	}()
	v := <-ch
	fmt.Println(v)
}
