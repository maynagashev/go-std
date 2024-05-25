package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	const n = 5

	for i := 0; i < n; i++ {
		wg.Add(1) // инкрементируем счётчик, сколько горутин нужно подождать

		go func(i int) {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("hi %d\n", i)
			// уменьшаем счётчик, когда горутина завершает работу
			wg.Done()
		}(i)
	}

	wg.Wait() // ждём все горутины
	fmt.Println("Всё готово")
}
