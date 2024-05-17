package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// TimeError предназначен для ошибок с фиксацией времени возникновения.
type TimeError struct {
	Time time.Time
	Err  error
}

// Error добавляет поддержку интерфейса error для типа TimeError.
func (te *TimeError) Error() string {
	return fmt.Sprintf("%v %v", te.Time.Format("2006/01/02 15:04:05"), te.Err)
}

// Unwrap возвращает исходную ошибку.
func (te *TimeError) Unwrap() error {
	return te.Err
}

// NewTimeError записывает ошибку err в тип TimeError c текущим временем.
func NewTimeError(err error) error {
	return &TimeError{
		Time: time.Now(),
		Err:  err,
	}
}

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
		// можем узнать оригинальную ошибку для TimeError
		fmt.Println("Original error:", errors.Unwrap(err))
		os.Exit(0)
	}
	// ...
}
