package lesson3

import (
	"fmt"
	"testing"
)

func InArray[T comparable](list []T, value T) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

//func InArrayInt(list []int, value int) bool {
//	for _, v := range list {
//		if v == value {
//			return true
//		}
//	}
//	return false
//}

func TestInArray(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	if !InArray(list, 3) {
		t.Error("not found")
	} else {
		t.Log("found")
	}
}

func TestInArrayString(t *testing.T) {
	list := []string{"a", "b", "c", "d", "e"}
	if !InArray(list, "c") {
		t.Error("not found")
	} else {
		t.Log("found")
	}
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

func Sort[T Number | string](list []T) {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}

func TestSort(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		list := []int16{5, 3, 1, 2, 4}
		Sort(list)
		t.Log(list)
	})
	t.Run("string", func(t *testing.T) {
		list := []string{"d", "a", "c", "b"}
		Sort(list)
		t.Log(list)
	})
}

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() T {
	if len(s.data) == 0 {
		var zero T
		return zero
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func TestStack(t *testing.T) {
	s := &Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	t.Log(s.Pop())
}

type AnyMap[T comparable] map[T]any

func IterateMap(m AnyMap[string]) {
	for k, v := range m {
		println(k)
		fmt.Printf("%T-%v\n", v, v)
	}
}

func TestIterateMap(t *testing.T) {
	m := AnyMap[string]{
		"name": "jack",
		"age":  18,
	}
	IterateMap(m)
}

func FilterSlice[T any](list []T, f func(T) bool) []T {
	var newList []T
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func TestFilterSlice(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	newList := FilterSlice(list, func(v int) bool {
		return v%2 == 0
	})
	t.Log(newList)
}
