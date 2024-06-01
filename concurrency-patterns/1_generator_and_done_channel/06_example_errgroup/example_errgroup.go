package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

type Name string

var names = []Name{"Anna", "Ivan", "Fedor", "Katya", "Gleb"}

// Hello — метод типа Name.
func (n Name) Hello() error {
	fmt.Printf("Hello %v!\n", n)
	return nil
}

func main() {
	g := &errgroup.Group{}

	for _, name := range names {
		// вызываем g.Go с method value в качестве аргумента
		g.Go(name.Hello)
	}

	/*
	   Ожидает завершения всех горутин в группе, как с ошибками, так и без ошибок.
	   - Команда errgroup.Group.Wait() - позволяет дождаться завершения всех горутин и получить информацию о возникших ошибках.
	   - Этот подход к обработке ошибок наиболее удобен и безопасен.
	*/
	g.Wait()
}
