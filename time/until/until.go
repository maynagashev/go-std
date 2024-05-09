package main

import (
	"fmt"
	"time"
)

// Разработчик Андрей родился 26 ноября 1993 года. Посчитайте количество дней до его 100-летия — относительно сегодняшнего дня.
func main() {

	birthday := time.Date(1993, time.November, 26, 0, 0, 0, 0, time.Local)
	hundred := birthday.AddDate(100, 0, 0)

	var days int
	//days = int(time.Until(hundred) / time.Hour / 24)
	days = int(time.Until(hundred).Hours() / 24)

	//days
	fmt.Println(birthday, hundred)
	fmt.Println("days: ", days)

	fmt.Println("Month", time.January.String())
}
