package main

import (
	"fmt"
	"sync"
	"time"
)

func startWorkers(c *sync.Cond, val *int) {
	workerCount := 3

	for i := 0; i < workerCount; i++ {
		go func(workerId int) {
			c.L.Lock()
			for {
				c.Wait()
				// получили сигнал
				fmt.Printf("val %v processed by worker %v\n", *val, workerId)
			}
		}(i)
	}
}

/*
set val to 0
val 0 processed by worker 0
val 0 processed by worker 2
val 0 processed by worker 1
set val to 1
val 1 processed by worker 1
val 1 processed by worker 0
val 1 processed by worker 2
set val to 2
val 2 processed by worker 2
val 2 processed by worker 1
val 2 processed by worker 0
set val to 3
val 3 processed by worker 0
val 3 processed by worker 1
val 3 processed by worker 2
*/
func main() {
	var m sync.Mutex
	c := sync.NewCond(&m)
	val := 0
	startWorkers(c, &val)
	// ждём, чтобы стартанули все воркеры
	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 4; i++ {
		m.Lock()
		val = i
		fmt.Printf("set val to %v\n", val)
		// отправляем сигнал всем воркерам
		c.Broadcast()
		m.Unlock()
		time.Sleep(time.Millisecond)
	}
}
