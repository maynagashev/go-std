package main

import (
	"errors"
	"log"

	"golang.org/x/sync/errgroup"
)

func main() {
	// создаём переменную errgroup
	g := new(errgroup.Group)

	// наши данные
	input := []int{1, 2, 3, 4}

	// генератор возвращает канал, через который он отправляет данные
	inputCh := generator(input)

	for data := range inputCh {
		// тут объявляем новую переменную внутри цикла, чтобы копировать переменную
		// в замыкание каждой горутины, а не использовать одно общее на всех значение.
		data := data

		// потребитель должен возвращать ошибку.
		// сигнатура анонимной функции всегда такая как в примере.
		g.Go(func() error {
			// получаем ошибку
			err := callDatabase(data)
			if err != nil {
				// возвращаем ошибку
				return err
			}

			return nil
		})
	}

	// здесь ждём выполнения горутин, и если хотя бы в одной из них возникает ошибка,
	// то присваиваем её err и обрабатываем. В этом случае просто выводим на экран.
	// Обратите внимание, что g.Wait() ждёт завершения всех запущенных горутин, даже
	// если приозошла ошибка.
	if err := g.Wait(); err != nil {
		log.Println(err)
	}
}

// generator возвращает канал, а затем отправляет в него данные
func generator(input []int) chan int {
	// создаём канал данных
	inputCh := make(chan int)

	// вызываем горутину в которой отправляем данные в канал inputCh
	go func() {
		// по завершении горутины закрываем канал
		defer close(inputCh)

		// перебираем данные в слайсе
		for _, data := range input {
			// отправляем данные из слайса в канал
			inputCh <- data
		}
	}()

	// возвращаем канал с данными
	return inputCh
}

// callDatabase просто возвращает ошибку
func callDatabase(data int) error {
	// допустим ошибка возникнет когда data = 3
	if data == 3 {
		return errors.New("ошибка запроса к базе данных")
	}

	return nil
}
