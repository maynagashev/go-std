package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// LabelError описывает ошибку с дополнительной меткой.
type LabelError struct {
	Label string // метка должна быть в верхнем регистре
	Err   error
}

// добавьте методы Error() и NewLabelError(label string, err error)
func (le *LabelError) Error() string {
	return fmt.Sprintf("[%s] %s", le.Label, le.Err)
}

// Unwrap требуется, в том числе чтобы errors.Is и errors.As работали и видели исходную ошибку.
func (le *LabelError) Unwrap() error {
	return le.Err
}

func NewLabelError(label string, err error) error {
	return &LabelError{
		Label: strings.ToUpper(label),
		Err:   err,
	}
}

func main() {
	_, err := os.ReadFile("mytest.txt")
	if err != nil {
		err = NewLabelError("file", err)
	}
	fmt.Println(errors.Is(err, os.ErrNotExist), err)
	// должна выводить
	// true [FILE] open mytest.txt: no such file or directory
}
