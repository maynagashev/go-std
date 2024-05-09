package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	currentTimeStr := "2021-09-19T15:59:41+03:00"
	currentTime, err := time.Parse(time.RFC3339, currentTimeStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currentTime)

	now := time.Now()

	// Проверка before
	fmt.Println("Is", now, "before", currentTime, "? Answer:", now.Before(currentTime))
	fmt.Println("Is", now, "equal", currentTime, "? Answer:", now.Equal(currentTime))
	fmt.Println("Is", now, "after", currentTime, "? Answer:", now.After(currentTime))

}
