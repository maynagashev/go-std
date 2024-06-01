package main

import (
	"fmt"
	"sync"
)

func main() {
	inCh := gen(2, 3)
	ch1 := square(inCh)
	ch2 := square(inCh)
	for n := range fanIn(ch1, ch2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for _, n := range nums {
			outCh <- n
		}
	}()

	return outCh
}

func square(inCh chan int) chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for n := range inCh {
			outCh <- n * n
		}
	}()

	return outCh
}

// fanIn принимает несколько каналов, в которых итоговые значения
func fanIn(chs ...chan int) chan int {
	var wg sync.WaitGroup
	outCh := make(chan int)

	// определяем функцию output для каждого канала в chs
	// функция output копирует значения из канала с в канал outCh, пока с не будет закрыт
	output := func(c chan int) {
		for n := range c {
			outCh <- n
		}
		wg.Done()
	}

	// добавляем в группу столько горутин, сколько каналов пришло в fanIn
	wg.Add(len(chs))
	// перебираем все каналы, которые пришли и отправляем каждый в отдельную горутину
	for _, c := range chs {
		go output(c)
	}

	// запускаем горутину для закрытия outCh после того, как все горутины отработают
	go func() {
		wg.Wait()
		close(outCh)
	}()

	// возвращаем общий канал
	return outCh
}
