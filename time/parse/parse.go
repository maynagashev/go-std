package main

import (
	"fmt"
	"time"
)

func main() {
	currentTimeStr := "2021-09-19T15:59:41+03:00"

	t, err := time.Parse(time.RFC3339, currentTimeStr)
	// или
	// currentTime, err := time.Parse("2006-01-02T15:04:05Z07:00", currentTimeStr)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("input: %s\nparsed: %#v\nvalue: %s\nlocal: %s", currentTimeStr, t, t.String(), t.Local())

	details()
}

func details() {
	now := time.Now()
	fmt.Println("Год:", now.Year())
	fmt.Println("Месяц:", now.Month())
	fmt.Println("Число:", now.Day())
	fmt.Println("День недели:", now.Weekday())
	hour, minutes, sec := now.Clock()
	fmt.Printf("Время: %d:%d:%d\n", hour, minutes, sec)
	fmt.Println("Часовой пояс:", now.Location())
	fmt.Println("timestamp в секундах:", now.Unix())
	fmt.Println("timestamp в наносекундах:", now.UnixNano())

	/*
		Месяц: May
		Число: 9
		День недели: Thursday
		Время: 9:20:36
		Часовой пояс: Local
		timestamp в секундах: 1715221236
		timestamp в наносекундах: 1715221236837525000
	*/
}
