package main

import (
	"fmt"
	"time"
)

// fatal error: concurrent map writes
func main() {
	m := make(map[int]int)

	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				//
				if _, ok := m[j]; !ok {
					m[j] = j
				}
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(len(m))
}
