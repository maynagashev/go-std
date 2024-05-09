package main

import (
	"fmt"
	"time"
)

// Напишите программу, которая 10 раз с интервалом в две секунды выведет разницу в секундах
// между текущим временем и временем запуска программы
func main() {

	start := time.Now()
	ticker := time.NewTicker(2 * time.Second)
	for i := range 10 {
		t := <-ticker.C
		fmt.Println(i, int(t.Sub(start).Seconds()))
	}
}
