/*
В языке Go вместо паттерна Строитель чаще всего применяют паттерн Функциональные опции.
Этот шаблон помогает решить те же задачи, но в полной мере использует возможности языка.

Суть паттерна заключается в том,
что для параметров создаются функции с замыканием,
которые в свою очередь возвращают функции,
принимающие объект и присваивающие ему нужный параметр.
*/
package main

import "fmt"

// Object — объект с параметром.
type Object struct {
	// данные объекта
	// ...
	// настраиваемые поля объекта
	Mode int
	Path string
}

// WithMode — пример функции, которая присваивает поле Mode.
func WithMode(mode int) func(*Object) {
	return func(o *Object) {
		o.Mode = mode
	}
}

// WithPath — пример функции, которая присваивает поле Path.
func WithPath(path string) func(*Object) {
	return func(o *Object) {
		o.Path = path
	}
}

// NewObject — функция-конструктор объекта.
func NewObject(opts ...func(*Object)) *Object {
	o := &Object{}

	// вызываем все указанные функции для установки параметров
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func main() {
	o := NewObject(WithMode(10), WithPath(`root`))
	fmt.Println(o)
}
