package test

import (
	"fmt"
	"testing"
)

func TestPtr(t *testing.T) {
	var a *int
	var s *string
	s1 := "abc"
	s = &s1
	fmt.Println(a, s, &s1)
}

func TestPtr2(t *testing.T) {
	i := new(int)
	fmt.Println(i)
	fmt.Printf("%T", i)
}

func TestPtr3(t *testing.T) {
	value := 10
	pointer := &value
	fmt.Println(value, pointer)
	fmt.Println(*pointer)
	*pointer = 20
	fmt.Println(*pointer, pointer)
	fmt.Println(value)
}
