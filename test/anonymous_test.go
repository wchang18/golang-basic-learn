package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestAFunc(t *testing.T) {
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(5, 10))

	var sub func(int, int) int
	sub = func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(15, 10))
}

func outer(num int) func() {
	count := 0
	inner := func() {
		count += num
		fmt.Println(count)
	}
	return inner
}

func TestClosure(t *testing.T) {
	counter := outer(2)
	counter()
	counter()
	counter()
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(s string) string {
		if !strings.HasSuffix(s, suffix) {
			return s + suffix
		}
		return s
	}
}

func TestSuffix(t *testing.T) {
	jpgFunc := makeSuffixFunc(".jpg")
	fmt.Println(jpgFunc("1"))
	fmt.Println(jpgFunc("2.jpg"))
}
