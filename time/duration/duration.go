package main

import (
	"fmt"
	"time"
)

// Для определения интервала времени используют тип Duration, который содержит количество наносекунд.
// Это переопределение типа int64.
func main() {
	round()

	truncate()
}

// Округлите текущее время до начала дня (полуночи), используя метод Truncate().
// Обратите внимание на время и часовой пояс в выводе.
func truncate() {
	var today time.Time
	today = time.Now().Truncate(24 * time.Hour)
	fmt.Println("truncated today", today)
	// truncated today 2024-05-09 07:00:00 +0700 +07
}

func round() {
	now := time.Now()
	// добавим 20 секунд к текущему времени
	fmt.Println(now.Add(20 * time.Second))
	// округлим время до часа
	fmt.Println(now.Round(time.Hour))
	// округлим время в меньшую сторону до начала трёхминутного интервала
	fmt.Println(now.Truncate(3 * time.Minute))

	/*
		При округлении времени интервалы отсчитываются от начала Unix-эпохи — это полночь 1 января 1970 года в UTC.
		Например, для 24-часовых форматов интервалы будут округляться до полуночи по UTC, а не до полуночи в локальном часовом поясе.
	*/
	// Truncate округляет время в меньшую сторону до начала
	// ближайшего указанного интервала.
	trunc := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local).Truncate(24 * time.Hour)
	fmt.Println(trunc)
}
