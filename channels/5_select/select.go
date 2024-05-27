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

//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	ch3 := make(chan int)
//
//	select {
//	case x := <-ch1:
//		// сценарий выполнится, если быстрее всего новое значение окажется в канале ch1
//	case y := <-ch2:
//		// сценарий выполнится, если быстрее всего новое значение окажется в канале ch2
//	case ch3 <- z:
//		// сценарий выполнится, если быстрее отправим значение в канал ch3
//	}
//}
