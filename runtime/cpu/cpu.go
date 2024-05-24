package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
Macbook Ядер: 12
Логических процессоров: 12 Горутин: 1
Логических процессоров: 2 Горутин: 2
*/
func main() {
	fmt.Println("Ядер:", runtime.NumCPU())

	// GOMAXPROCS(n int) int — изменяет значение GOMAXPROCS и возвращает количество логических процессоров, которое было установлено до вызова функции
	fmt.Println("Логических процессоров:", runtime.GOMAXPROCS(2),
		"Горутин:", runtime.NumGoroutine())
	go func() {
		time.Sleep(100 * time.Millisecond)
	}()
	fmt.Println("Логических процессоров:", runtime.GOMAXPROCS(0),
		"Горутин:", runtime.NumGoroutine())
}
