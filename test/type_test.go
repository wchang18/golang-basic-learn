package test

import (
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	type NewInt int  //定义新类型
	type MyInt = int //类型别名
	var a int
	var b MyInt = 2
	var c NewInt = 3
	a = b         //类型相同可以赋值
	c = NewInt(b) //类型不同不能赋值
	fmt.Println(a, b, c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
