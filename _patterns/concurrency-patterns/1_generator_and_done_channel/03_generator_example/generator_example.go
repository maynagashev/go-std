package main

import "fmt"

func main() {
	ch := generator("Hello")
	for msg := range ch {
		fmt.Println(msg)
	}
}

func generator(msg string) chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
		}
	}()

	return ch
}
