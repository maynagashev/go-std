package main

import "fmt"

func process(in1, in2 <-chan int, out chan<- int) {
loop:
	for {
		select {
		case x, ok := <-in1:
			if !ok {
				break loop
			}
			out <- x * 2
		case x, ok := <-in2:
			if !ok {
				break loop
			}
			out <- x * 3
		}
	}
	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	chout := make(chan int)

	go func() {
		for i := 0; i <= 20; i++ {
			// выбираем случайный канал для отправки значения
			select {
			case ch1 <- i:
			case ch2 <- i:
			}
		}
		close(ch1)
		close(ch2)
	}()
	go process(ch1, ch2, chout)
	for i := range chout {
		fmt.Printf("%d ", i)
	}
}
