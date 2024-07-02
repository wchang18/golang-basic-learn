package test

import (
	"fmt"
	"testing"
)

func TestConst(t *testing.T) {
	const i = 3.14
	fmt.Println(i)

	const (
		c1 = "abc"
		c4
		c2 = 20
		c3
	)
	fmt.Println(c1, c2, c3, c4)

	var a float64 = 20
	if a == c2 {
		fmt.Println("right")
	}
}

func TestIota(t *testing.T) {
	const (
		l1 = iota
		l2
		l3
	)
	fmt.Println(l1, l2, l3)

	const (
		_ = iota
		n1
		n2
		n3
		_
		n5
	)
	fmt.Println(n1, n2, n3, n5)
}

func TestIota2(t *testing.T) {

	const (
		_  = iota
		KB = 1 << (10 * iota) //二进制左移10位
		MB = 1 << (10 * iota) //二进制左移20位
		GB = 1 << (10 * iota) //二进制左移30位
	)

	const (
		a, b = iota, iota + 2
		c, d = iota, iota + 2
		e, f = iota, iota + 2
	)

	fmt.Println(KB, MB, GB)
	fmt.Println(a, b, c, d, e, f)
}
