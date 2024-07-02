package generics

import (
	"fmt"
	"testing"
)

// [T int | float64] 为类型形参
func Min[T int | float64](a, b T) T {
	if a > b {
		return b
	}
	return a
}

func TestMin(t *testing.T) {
	var (
		a float64 = 1
		b float64 = 2
		//c int32   = 5
		//d int32   = 6
	)
	fmt.Println(Min[float64](a, b))
	//fmt.Println(Min(c, d))
}

type Slice[T int | string] []T

type Map[K int | string, V float32 | int64] map[K]V

type Tree[T interface{}] struct {
	left, right *Tree[T]
	value       T
}

func TestMap(t *testing.T) {
	m1 := make(Map[int, float32])
	m1[1] = 1.1
	m1[2] = 2.2
	fmt.Printf("%#v\n", m1)

	l1 := make(Slice[string], 0)
	l1 = append(l1, "a")
	l1 = append(l1, "b")
	fmt.Printf("%#v\n", l1)

	t1 := Tree[int]{
		left:  &Tree[int]{},
		right: &Tree[int]{},
		value: 99,
	}
	fmt.Printf("%#v\n", t1)
}

type Signed interface {
	int | int8 | int16 | int32 | int64
}

type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Integer interface {
	Signed | Unsigned
}

type MyMap[Key Integer, Value any] map[Key]Value

func TestMyMap(t *testing.T) {
	m1 := MyMap[uint, bool]{
		1: true,
		2: false,
	}
	fmt.Printf("%#v\n", m1)
}
