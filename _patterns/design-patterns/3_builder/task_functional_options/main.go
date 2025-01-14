/*
На примере сборки спорткара и минивэна потренируйтесь применять шаблон Функциональные опции.

Реализуйте функции NewCar, WithSeats, WithDoors, WithABS.
*/
package main

import (
	"fmt"
)

type Car struct {
	Type  string
	Seats int // количество мест
	Doors int // количество дверей
	ABS   bool
}

// CarOptionFunc определяет тип функции для опций.
type CarOptionFunc func(*Car)

// NewCar создаёт автомобиль, используя указанные опции.
func NewCar(cartype string, opts ...CarOptionFunc) *Car {
	c := &Car{Type: cartype}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Car) String() string {
	return fmt.Sprintf("%s [seats:%d / doors: %d / abs: %t]",
		c.Type, c.Seats, c.Doors, c.ABS)
}

// WithSeats определяет количество мест в автомобиле.
func WithSeats(seats int) CarOptionFunc {
	return func(c *Car) {
		c.Seats = seats
	}
}

// WithDoors определяет количество дверей.
func WithDoors(doors int) CarOptionFunc {
	return func(c *Car) {
		c.Doors = doors
	}
}

// WithABS указывает на наличие ABS в автомобиле.
func WithABS() CarOptionFunc {
	return func(c *Car) {
		c.ABS = true
	}
}

func main() {
	// Можно добавить примеры использования здесь
	sportsCar := NewCar("sports car", WithSeats(2), WithDoors(2))
	minivan := NewCar("minivan", WithSeats(7), WithDoors(5), WithABS())

	fmt.Println(sportsCar)
	fmt.Println(minivan)
}
