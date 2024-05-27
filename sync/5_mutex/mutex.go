package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Неоптимальное решение с обычным мьютексом, часть горутин читают часть пишут.
func main() {
	var m sync.Mutex
	cache := map[int]int{}

	// горутины, которые изменяют мапу
	for i := 0; i < 10; i++ {
		go func() {
			for {
				m.Lock()
				cache[rand.Intn(5)] = rand.Intn(100)
				m.Unlock()
				time.Sleep(time.Second / 20)
			}
		}()
	}

	// горутины, которые читают мапу
	for i := 0; i < 10; i++ {
		go func() {
			for {
				m.Lock()
				fmt.Printf("%#v\n", cache)
				m.Unlock()
				time.Sleep(time.Second / 100)
			}
		}()
	}

	time.Sleep(1 * time.Second)
}
