package test

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	//for i := 1; i <= 5; i++ {
	//	fmt.Print(i)
	//	fmt.Print(",")
	//}
	//fmt.Println()

	//var n int
	//for ; n < 5; n++ {
	//	fmt.Print(n)
	//	fmt.Print(",")
	//}
	//fmt.Println()
	//
	var m int
	for m < 5 {
		fmt.Print(m)
		fmt.Print(",")
		m++
	}
	fmt.Println()
}

func TestForRange(t *testing.T) {
	//遍历数组
	//list := []int{3, 6, 9}
	//for i, item := range list {
	//	fmt.Println("index:", i, "item:", item)
	//}
	//
	//for i := range list {
	//	fmt.Println("index:", i)
	//}
	//
	//for _, item := range list {
	//	fmt.Println("item:", item)
	//}

	m := map[string]string{
		"name":    "tom",
		"address": "beijing",
	}
	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}
}

func TestGoto(t *testing.T) {
	for i := 0; i < 10; i++ {
		if i == 6 {
			goto here
		}
		fmt.Println(i)
	}

here:
	fmt.Println("调转到这里")
}

func TestBreak(t *testing.T) {
	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
}

func TestContinue(t *testing.T) {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
