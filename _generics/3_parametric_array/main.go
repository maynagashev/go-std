// Package main реализует функцию Map(), работающую для всех коллекций:
// массивов, слайсов, ассоциативных массивов, списков, деревьев и пр.
// Для этого применим паттерн Итератор.
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// comparable — новое ключевое слово в синтаксисе языка,
// оно обозначает типы, для которых определены операции `==` и `!=`
type Map[K comparable, V any] map[K]V
type Slice[T any] []T

type Slaps[K comparable, V any] interface {
	Map[K, V] | Slice[V]
}

func SuperMap[T Slaps[int, any]](s T, f func(T) T) T {
	for k, v := range s {
		// ошибка: cannot range over s
		// (variable of type T constrained by Slaps[int, any])
		// (T has no structural type)
		s[k] = f(v)
	}
	return s
}

type Iterator[V any] interface {
	First() (V, bool)
	Next() (V, bool)
	Set(V)
}

func Map[I Iterator[V], V any](it I, f func(V) V) {
	for item, ok := it.First(); ok; item, ok = it.Next() {
		it.Set(f(item))
	}
}

// Item — элемент списка со ссылкой на следующий элемент.
type Item[T any] struct {
	next  *Item[T]
	value T
}

// List — список.
type List[T any] struct {
	first *Item[T] // первый элемент
	cur   *Item[T] // текущий элемент
}

func NewList[V any](l *Item[V]) *List[V] {
	var i List[V]
	i.first = l
	i.cur = l
	return &i
}

func (l *List[T]) Next() (T, bool) {
	if l.cur.next != nil {
		l.cur = l.cur.next
		return l.cur.value, true
	}
	var empty T
	return empty, false
}

func (l *List[T]) First() (T, bool) {
	l.cur = l.first
	if l.cur == nil {
		var empty T
		return empty, false
	}
	return l.cur.value, true
}

func (l *List[T]) Set(v T) {
	if l.cur != nil {
		l.cur.value = v
	}
}

type MapGeneric[K constraints.Ordered, V any] map[K]V

type MapIter[K constraints.Ordered, V any] struct {
	m     MapGeneric[K, V]
	index []K
	cur   int
}

func NewMapIter[K constraints.Ordered, V any](m MapGeneric[K, V]) *MapIter[K, V] {
	var ret MapIter[K, V]
	ret.m = m
	ret.index = make([]K, 0, len(m))
	for k := range m {
		ret.index = append(ret.index, k)
	}
	return &ret
}

func (m *MapIter[K, V]) Next() (V, bool) {
	if m.cur < len(m.index)-1 {
		m.cur++
		return m.m[m.index[m.cur]], true
	}
	var empty V
	return empty, false
}

func (m *MapIter[K, V]) Set(v V) {
	if m.cur < len(m.index) {
		m.m[m.index[m.cur]] = v
	}
}

func (m *MapIter[K, V]) First() (V, bool) {
	m.cur = 0
	if len(m.index) > 0 {
		return m.m[m.index[m.cur]], true
	}
	var empty V
	return empty, false
}

func main() {
	m := map[string]string{
		"1": "a",
		"2": "b",
		"3": "c",
		"4": "d",
		"5": "e",
	}
	// конструируем итератор для мапы
	var iter = NewMapIter(m)

	for item, ok := iter.First(); ok; item, ok = iter.Next() {
		fmt.Println(item)
	}
	// применяем метод Map()
	Map(iter, Double[string])
	fmt.Println("After Mapping")

	for item, ok := iter.First(); ok; item, ok = iter.Next() {
		fmt.Println(item)
	}
}
