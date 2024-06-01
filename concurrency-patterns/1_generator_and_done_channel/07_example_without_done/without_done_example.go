package main

import (
	"fmt"
)

/*
В приведённом выше примере функция generator() заблокирована на неопределенный срок, поскольку получатель
не может получать данные после выхода из хендлера.
Завершить все горутины, которые больше не нужны, поможет только явная отмена.

Отправляем данные в канал: 1
Отправляем данные в канал: 2
Выходим из handler
*/
func main() {
	handler()
}

// handler получает данные из канала
func handler() {
	// слайс данных
	input := []int{1, 2, 3, 4, 5, 6}
	// получаем канал в который будут приходить данные из generator
	inputCh := generator(input)

	// перебираем данные из канала
	for data := range inputCh {
		// если среди данных из канала встретится 1, то выходим из handler
		if data == 1 {
			fmt.Println("Выходим из handler")
			return
		}
	}
}

// generator отправляет данные из слайса в канал, а потом его возвращает
func generator(input []int) chan int {
	inputCh := make(chan int)

	go func() {
		// по завершении закрываем канал inputCh
		defer close(inputCh)

		// передаём данные в канал inputCh
		for _, data := range input {
			fmt.Printf("Отправляем данные в канал: %d\n", data)
			inputCh <- data
		}
	}()

	// возвращаем канал inputCh
	return inputCh
}
