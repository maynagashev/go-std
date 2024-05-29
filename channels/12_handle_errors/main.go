package main

import (
	"errors"
	"log"
)

// структура, в которую добавили ошибку
type Result struct {
	data int
	err  error
}

func main() {
	// ваши данные
	input := []int{1, 2, 3, 4}

	// канал с результатами работы функции consumer
	resultCh := make(chan Result)

	// получаем канал с данными из генератора
	inputCh := generator(input)

	// порождаем горутину которая отправляет результат в resultCh вместе с ошибкой
	go consumer(inputCh, resultCh)

	// читаем результаты из канала resultCh
	for res := range resultCh {
		if res.err != nil {
			// здесь обрабатываем ошибку как обычно в Go
			log.Println("разберемся с ошибкой здесь")
		}
	}
}

// consumer вызывает другую функцию, которая возвращает ошибку
func consumer(inputCh chan int, resultCh chan Result) {
	// закрваем resultCh при завершении функции consumer
	defer close(resultCh)

	// перебираем данные из канала inputCh
	for data := range inputCh {
		// получаем ошибку
		resp, err := callDatabase(data)

		// создаем структуру
		result := Result{
			data: resp,
			err:  err,
		}

		// отправляем структуру в канал
		resultCh <- result
	}
}

// generator отправляет данные в канал inputCh
func generator(input []int) chan int {
	// создаём канал, куда будем отправлять данные из слайса
	inputCh := make(chan int)

	// через отдельную горутину генератор отправляет данные в канал
	go func() {
		defer close(inputCh)

		// перебираем данные из слайса
		for _, data := range input {
			// отправляем данные в канал
			inputCh <- data
		}
	}()

	// возвращаем канал с данными
	return inputCh
}

// callDatabase просто возвращает ошибку как бы из функции обращения к базе данных
func callDatabase(data int) (int, error) {
	return data, errors.New("ошибка запроса к базе данных")
}
