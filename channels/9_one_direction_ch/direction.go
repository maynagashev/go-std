package main

import (
	"fmt"
	"strings"
)

// Generate отправляет в канал out односимвольные строки.
func Generate(out chan<- string) {
	for ch := 'a'; ch <= 'z'; ch++ {
		out <- string([]rune{ch})
	}
	close(out)
}

// Process читает строки из канала in, переводит их в верхний регистр
// и отправляет в канал out.
func Process(in <-chan string, out chan<- string) {
	for v := range in {
		out <- strings.ToUpper(v)
	}
	close(out)
}

func main() {
	lower := make(chan string)
	upper := make(chan string)
	go Generate(lower)
	go Process(lower, upper)

	// выводим строки из канала upper по мере получения
	for s := range upper {
		fmt.Print(s)
	}
}
