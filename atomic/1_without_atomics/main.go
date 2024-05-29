package main

import (
	"fmt"
	"sync"
)

func count() {
	var counter int64

	var wg sync.WaitGroup

	// 25 горутин увеличивают значение счётчика
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 2000; i++ {
				counter++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("%d ", counter)
}

func main() {
	// делаем несколько попыток
	for i := 0; i < 5; i++ {
		count()
	}
}
