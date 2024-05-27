package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// меняем тип мьютекса
	var m sync.RWMutex
	cache := map[int]int{}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				// здесь остаются блокировки на запись
				m.Lock()
				cache[rand.Intn(5)] = rand.Intn(100)
				m.Unlock()
				time.Sleep(time.Second / 20)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				// при чтении используем Rlock() и RUnlock()
				m.RLock()
				fmt.Printf("%#v\n", cache)
				m.RUnlock()
				time.Sleep(time.Second / 100)
			}
		}()
	}

	time.Sleep(1 * time.Second)
}
