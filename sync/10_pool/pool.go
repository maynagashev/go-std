package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Создание нового объекта")
			return 0
		},
	}

	// Кладем несколько объектов в пул
	pool.Put(10)
	pool.Put(20)
	pool.Put(30)

	// Достаем объекты (порядок не гарантируется)
	obj1 := pool.Get().(int)
	fmt.Println("obj1:", obj1)

	obj2 := pool.Get().(int)
	fmt.Println("obj2:", obj2)

	obj3 := pool.Get().(int)
	fmt.Println("obj3:", obj3)

	// Если запрашиваем больше объектов, чем положили, создастся новый (0)
	obj4 := pool.Get().(int)
	fmt.Println("obj4:", obj4)
}
