package main

import (
	"fmt"
	"os"
	"time"
)

func ReadTextFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return ``, NewTimeError(err)
	}
	return string(data), nil
}

func main() {
	_, err := ReadTextFile("myconfig.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// ...
}

// TimeError предназначен для ошибок с фиксацией времени возникновения.
type TimeError struct {
	Time time.Time
	Err  error
}

// Error добавляет поддержку интерфейса error для типа TimeError.
func (te *TimeError) Error() string {
	return fmt.Sprintf("%v %v", te.Time.Format("2006/01/02 15:04:05"), te.Err)
}

// NewTimeError записывает ошибку err в тип TimeError c текущим временем.
func NewTimeError(err error) error {
	return &TimeError{
		Time: time.Now(),
		Err:  err,
	}
}
