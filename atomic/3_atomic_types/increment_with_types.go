package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func count() {
	var counter atomic.Int64

	var wg sync.WaitGroup

	// горутины увеличивают значение счётчика
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 2000; i++ {
				counter.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("%d ", counter.Load())
}

func main() {
	// делаем несколько попыток
	for i := 0; i < 5; i++ {
		count()
	}
}
