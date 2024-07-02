package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestVar(t *testing.T) {
	//单个声明
	var name string

	//批量声明
	var (
		i       int64
		age     uint8
		money   float64
		isAdmin bool
	)

	//声明加初始化
	var val int32 = 12

	//快捷声明，默认声明为string
	key := "apple"

	fmt.Println(name, i, age, money, isAdmin, val, key)
}

func TestVar1(t *testing.T) {
	str := "dog"
	num := 100
	flt := 9.99
	ok := true

	fmt.Println(reflect.TypeOf(str)) //string
	fmt.Println(reflect.TypeOf(num)) //int
	fmt.Println(reflect.TypeOf(flt)) //float64
	fmt.Println(reflect.TypeOf(ok))  //bool
}

func TestVar2(t *testing.T) {
	name, _ := getUser()
	_, age := getUser()
	fmt.Println(name, age)
}

func getUser() (string, int) {
	return "tom", 20
}
