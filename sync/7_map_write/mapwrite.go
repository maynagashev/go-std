package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := make(map[int]int)
	var mutex = sync.Mutex{}

	for i := 0; i < 100; i++ {
		go func(v int) {
			mutex.Lock()
			m[v] = 1
			mutex.Unlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(len(m))
}
