package main

import "fmt"

func main() {
	input := 1

	// пример 1
	// Вызываем функцию сложения, результат которой становится аргументом первого параметра функции умножения
	fmt.Println(multiply(add(input, 1), 2))

	// пример 2
	// можно поменять местами этапы, чтобы получить другой результат
	fmt.Println(add(multiply(input, 1), 2))
}

// add функция сложения
func add(x int, y int) int {
	return x + y
}

// multiply функция умножения
func multiply(x int, y int) int {
	return x * y
}
