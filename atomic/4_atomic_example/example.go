package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter atomic.Int64

func worker(wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		counter.Add(1)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	// программа должна выводить 200000
	fmt.Println(counter.Load())
}
