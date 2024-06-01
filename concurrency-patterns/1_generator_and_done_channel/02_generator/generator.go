package main

import "fmt"

func main() {
	// данные в слайсе, которые будем отправлять
	input := []int{1, 2, 3, 4, 5, 6}

	// получаем канал с данными из генератора
	inputCh := generator(input)

	// отправляем данные потребителю через канал inputCh
	consumer(inputCh)
}

// generator — генератор, который создает канал и сразу возвращает его
func generator(input []int) chan int {
	inputCh := make(chan int)

	// через отдельную горутину генератор отправляет данные в канал
	go func() {
		// закрываем канал по завершению горутины — это отправитель
		defer close(inputCh)

		// перебираем данные в слайсе
		for _, data := range input {
			// отправляем данные в канал inputCh
			inputCh <- data
		}
	}()

	// возвращаем канал inputCh
	return inputCh
}

// consumer — потребитель проходит через канал и одновременно обрабатывает
// данные из него (выводит на экран)
func consumer(inputCh chan int) {
	for data := range inputCh {
		fmt.Println(data)
	}
}
