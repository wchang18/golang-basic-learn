package test

import (
	"fmt"
	"testing"
)

func deferFunc() {
	defer Print("结束")

	fmt.Println("开始")
}

func deferFunc2() {
	defer Print("结束1")
	defer Print("结束2")
	defer Print("结束3")

	fmt.Println("开始")
}

func Print(v string) {
	fmt.Println(v)
}

func TestDefer(t *testing.T) {
	//deferFunc()
	deferFunc2()
}

func TestPanic(t *testing.T) {
	defer Print("close")
	fmt.Println("开始")
	panic("退出")
	fmt.Println("结束")
}

func TestRecover(t *testing.T) {
	defer func() {
		if res := recover(); res != nil {
			fmt.Println("panic内容：", res)
		}
	}()

	panic("有错误，退出")
	fmt.Println("结束")
}

func TestNew(t *testing.T) {
	v := new(int)
	fmt.Println(v)

	s := make([]int, 0)
	fmt.Println(s)
}
