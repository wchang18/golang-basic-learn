package main

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	fmt.Println("hello world")
}

func TestTwo(t *testing.T) {
	fmt.Println("hello2")
	for i := range 5 {
		println(i)
	}
}
