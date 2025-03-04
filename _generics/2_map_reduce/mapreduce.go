package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Slice[T any] []T

func (s *Slice[T]) Map(f func(T) T) *Slice[T] {
	for k, v := range *s {
		(*s)[k] = f(v)
	}
	return s
}

func (s *Slice[T]) Reduce(r T, f func(a, e T) T) T {
	for _, v := range *s {
		r = f(r, v)
	}
	return r
}

func (s *Slice[T]) Filter(allow func(e T) bool) *Slice[T] {
	var result Slice[T]
	for _, v := range *s {
		if allow(v) {
			result = append(result, v)
		}
	}
	*s = result
	return s
}

func Sum[T constraints.Ordered](a, e T) T {
	return a + e
}

func Double[T constraints.Ordered](v T) T {
	return v + v
}

func isEven(i int) bool {
	return i%2 == 0
}

func main() {
	var si = Slice[int]{1, 2, 3, 4, 5}
	sum := si.Reduce(0, Sum[int])
	fmt.Println(sum)

	// теперь цепочку
	res := si.Map(Double[int]).Reduce(0, Sum[int])
	fmt.Println(res)

	// теперь для строк
	var ss = Slice[string]{"foo", "bar", "buzz"}
	res1 := ss.Map(Double[string]).Reduce("", Sum[string])
	fmt.Println(res1)

	var si = Slice[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res2 := si.Filter(isEven).Map(Double[int]).Reduce(0, Sum[int])
	fmt.Println(res2)
}
