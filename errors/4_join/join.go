package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Settings struct {
	Host string
	Port int
}

var (
	ErrNoHost = errors.New("Не указан host")
	ErrNoPort = errors.New("Не указан port")
)

func ParseSettings(input string) (*Settings, error) {
	var settings Settings
	err := json.Unmarshal([]byte(input), &settings)
	if err != nil {
		return nil, err
	}
	// находим сразу все ошибки
	var errs []error
	if len(settings.Host) == 0 {
		errs = append(errs, ErrNoHost)
	}
	if settings.Port == 0 {
		errs = append(errs, ErrNoPort)
	}
	return &settings, errors.Join(errs...)
}

func main() {
	settings, err := ParseSettings(`{"host":"localhost", "port": 3000}`)
	fmt.Println(err, settings)

	_, err = ParseSettings("{}")
	fmt.Println("print:", err)
	fmt.Printf("err: %#v\n", err)
	fmt.Println(errors.Is(err, ErrNoHost), errors.Is(err, ErrNoPort))

	// unwrap ошибки после join возвращает null
	fmt.Printf("unwrap %#v", errors.Unwrap(err))
}
