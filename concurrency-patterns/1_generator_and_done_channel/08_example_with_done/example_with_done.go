package main

import (
	"log"
	"time"
)

/*
Стоп-кран.
Приём из закрытого канала возвращает нулевое значение.
Используя это свойство, можно создать канал Done и закрыть его для завершения всех горутин.
*/
func main() {
	input := []int{1, 2, 3, 4, 5, 6}
	handler(input)
	time.Sleep(time.Second)
}

// handler получает данные из слайса
func handler(input []int) {
	// канал для явной отмены
	doneCh := make(chan struct{})
	// когда выходим из handler — сразу закрываем канал doneCh
	defer close(doneCh)

	// теперь передаём и канал отмены doneCh
	inputCh := generator(doneCh, input)

	// забираем данные из канала
	for data := range inputCh {
		// если в данных 3 — выходим из handler
		if data == 3 {
			log.Println("Прекращаем обработку данных из канала")
			return
		}
		log.Println(data)
	}
	log.Println("Данные во входном канале закончились")
}

// generator возвращает канал с данными
func generator(doneCh chan struct{}, input []int) chan int {
	// канал, в который будем отправлять данные из слайса
	inputCh := make(chan int)

	// горутина, в которой отправляются данные в канал inputCh
	go func() {
		// по завершении закрываем канал inputCh
		defer close(inputCh)

		// перебираем данные в слайсе input
		for _, data := range input {
			select {
			// если канал doneCh закрылся - сразу выходим из горутины
			case <-doneCh:
				log.Println("Останавливаем генератор")
				return
			// отправляем данные в канал inputCh
			case inputCh <- data:
			}
		}
	}()

	// возвращаем канал с данными
	return inputCh
}
