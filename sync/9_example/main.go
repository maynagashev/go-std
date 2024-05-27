package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		go func(v int) {
			wg.Add(1)
			for j := 0; j < 10; j++ {
				mu.Lock()

				k := 10*v + j
				m[k] = k
				mu.Unlock()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(m))
}
