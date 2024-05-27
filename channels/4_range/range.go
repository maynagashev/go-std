package main

import "fmt"

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x // посылаем значения в канал
		x, y = y, x+y
	}

	close(ch) // закрываем канал
}

func main() {
	ch := make(chan int, 7)
	// специально делаем буфер канала меньше,
	// чем количество чисел Фибоначчи
	go fibonacci(15, ch)

	for i := range ch {
		// считываем значения из канала, пока он не будет закрыт
		fmt.Printf("%d ", i)
	}
}

func basic() {
	ch := make(chan int, 7)
	for i := range ch {
		fmt.Println(i)
	}
}

func same() {
	ch := make(chan int, 7)
	for {
		i, ok := <-ch
		if !ok {
			break
		} else {
			fmt.Println(i)
		}
		// тело цикла
	}
}
