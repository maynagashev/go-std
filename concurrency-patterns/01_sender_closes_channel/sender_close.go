package main

import "fmt"

func main() {
	// создаём канал
	ch := make(chan int)

	// вызываем горутину отправителя
	go sender(ch)

	// вызываем получателя
	recipient(ch)
}

// sender отправляет в канал числа от 0 до 9
func sender(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	// закрываем канал после отправки
	close(ch)
}

// recipient забирает из канала значения и выводит на экран,
// когда канал закрыт, выходит из цикла и завершает функцию
func recipient(ch chan int) {
	// читаем данные из канала, пока он открыт
	for data := range ch {
		// и выводим их на экран
		fmt.Println(data)
	}
}
