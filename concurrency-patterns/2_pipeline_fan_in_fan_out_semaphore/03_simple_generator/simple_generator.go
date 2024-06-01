package main

import "fmt"

func main() {
	c := gen(2, 3)
	out := square(c)

	for res := range out {
		fmt.Println(res)
	}
}

// реализация генератора gen здесь
func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

// реализация square здесь
func square(in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()

	return out
}
