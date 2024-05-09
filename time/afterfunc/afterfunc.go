package main

import (
	"fmt"
	"time"
)

func main() {
	time.AfterFunc(1*time.Second, func() {
		fmt.Println("Hi from AfterFunc")
	})
	fmt.Println("Hi")
	// ожидаем 2 секунды, чтобы успела запуститься функция в AfterFunc
	time.Sleep(2 * time.Second)
	fmt.Println("Goodbye")
}
