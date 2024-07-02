package test

import (
	"errors"
	"fmt"
	"testing"
)

// 不带参数
func SayHello() {
	fmt.Println("hello")
}

// 带一个参数
func Say(word string) {
	fmt.Println(word)
}

// 带多个参数
func SaySomeTimes(some int, word string) {
	for i := 0; i < some; i++ {
		fmt.Println(i, word)
	}
}

func TestSayFunc(t *testing.T) {
	SayHello()
	Say("你好")
	SaySomeTimes(3, "welcome")
}

// 一个返回值
func AddTwo(a, b int) int {
	return a + b
}

// 不定参数，可以为0个，也可以为多个
func Sum(number ...int) int {
	var res int
	for _, item := range number {
		res += item
	}
	return res
}

// 返回值的变量可以提前定义
func Sum2(number ...int) (res int) {
	for _, item := range number {
		res += item
	}
	return
}

// 返回多个值
func Check(i int) (bool, error) {
	if i > 10 {
		return false, errors.New("number not greater than 10")
	} else {
		return true, nil
	}
}

func TestAddFunc(t *testing.T) {
	fmt.Println(AddTwo(1, 2))
	fmt.Println(Sum(1, 2, 3))
	nums := []int{4, 5, 6}
	fmt.Println(Sum2(nums...))
	fmt.Println(Check(9))
	fmt.Println(Check(11))
}

var Hour = 3600

func ChangeHour() {
	fmt.Println(Hour)
	Hour = 1200
}

func TestGlobal(t *testing.T) {
	day := 10
	ChangeHour()
	fmt.Println(Hour)
	fmt.Println(day)
}
