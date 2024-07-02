package test

import (
	"fmt"
	"testing"
)

func TestBit2(t *testing.T) {
	b1 := 0b1001
	b2 := 0b1100
	fmt.Printf("%b\n", b1&b2)
	fmt.Printf("%b\n", b1|b2)
	fmt.Printf("%b\n", b1^b2)

	val := 0b10110
	mark := 0b01100
	fmt.Printf("%b\n", val&^mark)

	num := 16
	big := num << 2    //二进制像左移动两位，原数字增长4倍
	little := num >> 2 //二进制像右移动两位，原数字下降4倍
	fmt.Println(big, little)
}

func TestBit3(t *testing.T) {
	var a = 1
	for a < 10 {
		a += 2 // 相当于 a = a+2
	}
	fmt.Println(a)
}
