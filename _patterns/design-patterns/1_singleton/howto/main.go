package main

import (
	"fmt"
	"sync"
	"time"
)

type singleInstance struct {
}

var (
	singleton *singleInstance
	once      sync.Once
)

func getSingleton() *singleInstance {
	once.Do( // функция ниже выполнится только один раз
		func() {
			// инициализируем объект
			fmt.Println("Инициализируем singleton")
			singleton = &singleInstance{}
		})

	return singleton
}

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("Адрес singleton: %p\n", getSingleton())
		}(i)
	}
	time.Sleep(500 * time.Millisecond)
}
