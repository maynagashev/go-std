package main

import "fmt"

func main() {
	ch := generator("Hello")
	for msg := range ch {
		fmt.Println(msg)
	}
}

// Тут ваш генератор
/*
Hello 0
Hello 1
Hello 2
Hello 3
Hello 4
*/
func generator(msg string) chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
		}
		close(ch)
	}()
	return ch
}
