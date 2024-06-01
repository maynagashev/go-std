package main

import (
	"context"
	"log"
	"time"
)

/*
Остановить выполнение горутины, снабжённой контекстом, можно, прочитав в канале сигнал от контекста: <-ctx.Done().
ctx.Done() сигнализирует о завершении задачи.

Т.е. вместо отменяющего канала используем Context.Done()
*/
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	input := []int{1, 2, 3, 4, 5, 6}

	// Хэндлер запускается в горутине, после того как он отработает, отменяем контекст
	go func() {
		handler(ctx, input)
		cancel()
	}()

	time.Sleep(time.Second)
}

// передадим контекст и данные из слайса
func handler(ctx context.Context, input []int) {
	// передаём данные и контекст в генератор
	inputCh := generator(ctx, input)

	// теперь канал для отмены не нужен

	for data := range inputCh {
		if data == 3 {
			log.Println("Прекращаем обработку данных из канала")
			return
		}
		log.Println(data)
	}

	// до сюда не дойдём, т.к. генератор закроется раньше если в данных встретится 3
	log.Println("Данные во входном канале закончились")
}

func generator(ctx context.Context, input []int) chan int {
	inputCh := make(chan int)

	// генератор также запускается в горутине
	go func() {
		defer close(inputCh)

		for _, data := range input {
			select {
			// вместо отменяющего канала используем Context.Done()
			case <-ctx.Done():
				log.Println("Пришел сигнал ctx.Done(). Останавливаем генератор")
				return
			case inputCh <- data:
			}
		}
	}()

	return inputCh
}
