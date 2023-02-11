package main

import (
	"fmt"
	"strings"
)

func main() {
	// var s Stack
	// s.Push(10)
	// s.Push(20)
	// s.Push(30)
	// fmt.Println(s)
	// v, ok := s.Pop()
	// fmt.Println(v, ok)
	// fmt.Println(s)

	// var sg StackG[string]
	// sg.Push("I")
	// sg.Push("hate")
	// sg.Push("venturas")
	// sg.Push("LTD")
	// fmt.Println(sg)
	// val, ok := sg.Pop()
	// if ok {
	// 	fmt.Println(val, ok)
	// }
	// fmt.Println(sg)
	// fmt.Println(reflect.TypeOf(sg))
	// fmt.Println(sg.Contains("nope"))
	// fmt.Println(sg.Contains("love"))
	// fmt.Println(sg.Contains("hate"))

	var it *Tree[OrderableInt]
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(12))
	a := 10
	// uncomment to see a compile-time error
	// it = it.Insert(a)
	it = it.Insert(OrderableInt(a))
	// uncomment to see a compile-time error
	// it = it.Insert(OrderableInt("nope"))
}

// Map turns a []T1 to a []T2 using a mapping function.
// This function has two type parameters, T1 and T2.
// This works with slices of any type
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

type IntOrFloat interface {
	type int, float64
}
type Orderable[T any] interface {
	Order(T) int
}

type Tree[T Orderable[T]] struct {
	val         T
	left, right *Tree[T]
}

type OrderableInt int

func (oi OrderableInt) Order(val OrderableInt) int {
	return int(oi - val)
}

type OrderableString string

func (os OrderableString) Order(val OrderableString) int {
	return strings.Compare(string(os), string(val))
}

func (t *Tree[T]) Insert(val T) *Tree[T] {
	if t == nil {
		return &Tree[T]{val: val}
	}
	switch comp := val.Order(t.val); {
	case comp < 0:
		t.left = t.left.Insert(val)
	case comp > 0:
		t.right = t.right.Insert(val)
	}
	return t
}

func (t *Tree[T]) Contains(val T) bool {
	if t == nil {
		return false
	}
	switch comp := val.Order(t.val); {
	case comp < 0:
		return t.left.Contains(val)
	case comp > 0:
		return t.right.Contains(val)
	default:
		return true
	}
}

///////////////////////////////////////////////////////////////////////////////////////////

type Stack struct {
	vals []interface{}
}

func (s *Stack) Push(val interface{}) {
	s.vals = append(s.vals, val)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.vals) == 0 {
		return nil, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

type StackG[T comparable] struct {
	vals []T
}

func (s *StackG[T]) Push(val T) {
	s.vals = append(s.vals, val)
}
func (s *StackG[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s *StackG[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}
