package main

import (
	"fmt"
	"sync"
)

/*
8
1
5
6
7
2
3
0
4
fatal error: all goroutines are asleep - deadlock!

Когда программа засыпает и done меньше ожидаемых.
*/
func main() {
	var wg sync.WaitGroup

	n := 10
	wg.Add(n)
	for i := 0; i < n-1; i++ {
		go func(v int) {
			fmt.Println(v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
