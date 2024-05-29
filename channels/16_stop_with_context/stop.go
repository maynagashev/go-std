package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	input := []int{1, 2, 3, 4, 5, 6}

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
	log.Println("Данные во входном канале закончились")
}

func generator(ctx context.Context, input []int) chan int {
	inputCh := make(chan int)

	go func() {
		defer close(inputCh)

		for _, data := range input {
			select {
			// вместо отменяющего канала используем Context.Done()
			case <-ctx.Done():
				log.Println("Останавливаем генератор")
				return
			case inputCh <- data:
			}
		}
	}()

	return inputCh
}
