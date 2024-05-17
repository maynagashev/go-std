package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func ReadTextFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		// добавляем время и обёртываем ошибку
		now := time.Now().Format("2006/01/02 15:04:05")
		return "", fmt.Errorf("%s %w", now, err)
	}
	return string(data), nil
}

func main() {
	data, err := ReadTextFile("myconfig.yaml")
	if err != nil {
		fmt.Println(err)
		// можем узнать оригинальную ошибку
		fmt.Println("Original error:", errors.Unwrap(err))
		os.Exit(0)
	}
	fmt.Println(data)
	// ...
}
