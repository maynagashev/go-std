package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	timer := time.NewTimer(2 * time.Second) // создаём таймер
	t := <-timer.C                          // ожидаем срабатывания таймера, сообщение из канала
	fmt.Println(t.Sub(start).Seconds())     // выводим разницу во времени
}
