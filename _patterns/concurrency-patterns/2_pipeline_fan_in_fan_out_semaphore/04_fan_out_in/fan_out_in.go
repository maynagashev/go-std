/*
В паттерне Fan-Out создаётся большее количество горутин на определенном этапе, что увеличивает его пропускную способность.

Паттерн Fan-In объединяет несколько результатов в один канал.
Этот процесс также называют мультиплексированием.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// слайс данных
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	// сигнальный канал для завершения горутин
	doneCh := make(chan struct{})
	// закрываем его при завершении программы
	defer close(doneCh)

	// канал с данными
	inputCh := generator(doneCh, input)

	// получаем слайс каналов из 10 рабочих add
	channels := fanOut(doneCh, inputCh)

	// а теперь объединяем десять каналов в один
	addResultCh := fanIn(doneCh, channels...)

	// передаём тот один канал в следующий этап обработки
	resultCh := multiply(doneCh, addResultCh)

	// выводим результаты расчетов из канала
	for res := range resultCh {
		fmt.Println(res)
	}
}

// generator функция из предыдущего примера, делает то же, что и делала
func generator(doneCh chan struct{}, input []int) chan int {
	inputCh := make(chan int)

	go func() {
		defer close(inputCh)

		for _, data := range input {
			select {
			case <-doneCh:
				return
			case inputCh <- data:
			}
		}
	}()

	return inputCh
}

// multiply функция из предыдущего примера, делает то же, что и делала
func multiply(doneCh chan struct{}, inputCh chan int) chan int {
	multiplyRes := make(chan int)

	go func() {
		defer close(multiplyRes)

		for data := range inputCh {
			result := data * 2

			select {
			case <-doneCh:
				return
			case multiplyRes <- result:
			}
		}
	}()
	return multiplyRes
}

// add функция из предыдущего примера, делает то же, что и делала
func add(doneCh chan struct{}, inputCh chan int) chan int {
	addRes := make(chan int)

	go func() {
		defer close(addRes)

		for data := range inputCh {
			// замедлим вычисление, как будто функция add требует больше вычислительных ресурсов
			time.Sleep(time.Second)

			result := data + 1

			select {
			case <-doneCh:
				return
			case addRes <- result:
			}
		}
	}()
	return addRes
}

// fanOut принимает канал данных, порождает 10 горутин
func fanOut(doneCh chan struct{}, inputCh chan int) []chan int {
	// количество горутин add
	numWorkers := 10
	// каналы, в которые отправляются результаты
	channels := make([]chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		// получаем канал из горутины add
		addResultCh := add(doneCh, inputCh)
		// отправляем его в слайс каналов
		channels[i] = addResultCh
	}

	// возвращаем слайс каналов
	return channels
}

// fanIn объединяет несколько каналов resultChs в один.
func fanIn(doneCh chan struct{}, resultChs ...chan int) chan int {
	// конечный выходной канал в который отправляем данные из всех каналов из слайса, назовём его результирующим
	finalCh := make(chan int)

	// понадобится для ожидания всех горутин
	var wg sync.WaitGroup

	// перебираем все входящие каналы
	for _, ch := range resultChs {
		// в горутину передавать переменную цикла нельзя, поэтому делаем так
		chClosure := ch

		// инкрементируем счётчик горутин, которые нужно подождать
		wg.Add(1)

		go func() {
			// откладываем сообщение о том, что горутина завершилась
			defer wg.Done()

			// получаем данные из канала
			for data := range chClosure {
				select {
				// выходим из горутины, если канал закрылся
				case <-doneCh:
					return
				// если не закрылся, отправляем данные в конечный выходной канал
				case finalCh <- data:
				}
			}
		}()
	}

	go func() {
		// ждём завершения всех горутин
		wg.Wait()
		// когда все горутины завершились, закрываем результирующий канал
		close(finalCh)
	}()

	// возвращаем результирующий канал
	return finalCh
}
