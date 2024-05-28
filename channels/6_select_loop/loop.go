package main

import "fmt"

func fibonacci(ch chan int, quit chan bool) {
	x, y := 0, 1
loop: // метка цикла
	for {
		select {
		case ch <- x: // ждём, когда заберут значение из канала,
			// чтобы сгенерировать следующее
			x, y = y, x+y
		case <-quit: // параллельно ждём сигнала об окончании работы
			break loop
		}
	}
	fmt.Println("Выход")
}

func main() {
	// Канал используется для возврата значений и их печати в отдельной горутине
	ch := make(chan int)
	quit := make(chan bool)

	// запускаем горутину для печати 15 значений из канала и отправляем сигнал об окончании
	go func() {
		for i := 0; i < 15; i++ {
			fmt.Println(<-ch)
		}
		// подаём сигнал об окончании работы
		quit <- true
	}()

	fibonacci(ch, quit)
}
