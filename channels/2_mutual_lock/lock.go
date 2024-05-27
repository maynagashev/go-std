package main

import "fmt"

// fatal error: all goroutines are asleep - deadlock!
func main() {
	ch := make(chan int)
	// запись в небуферизованный канал никто не прочитает
	ch <- 10
	fmt.Println(<-ch)
}
