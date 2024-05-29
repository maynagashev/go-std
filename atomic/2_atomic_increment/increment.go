package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func count() {
	var counter int64

	var wg sync.WaitGroup

	// горутины увеличивают значение счётчика
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 2000; i++ {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("%d ", atomic.LoadInt64(&counter))
}

func main() {
	// делаем несколько попыток
	for i := 0; i < 5; i++ {
		count()
	}
}
