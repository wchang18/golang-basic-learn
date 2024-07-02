package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetType2(t *testing.T) {
	var f float64
	var i int64
	var s string
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(s))
}

func GetType(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Printf("type:%v, kind:%v\n", t.Name(), t.Kind())
}

type MyInt int

type Users struct {
	Name string
}

func TestGetType(t *testing.T) {
	var a int
	GetType(a)
	var b MyInt
	GetType(b)
	var c rune
	GetType(c)
	var u = Users{
		Name: "jack",
	}
	GetType(u)
	var l []string
	GetType(l)
}

func TestValue(t *testing.T) {

	var i int64 = 100
	SetValue(&i)
	fmt.Println(i)

	var a *int
	v := reflect.ValueOf(a)
	fmt.Println(v.IsValid())
	fmt.Println(v.IsNil())
	fmt.Println("a---")
	//
	var b float64
	b = 1
	v = reflect.ValueOf(b)
	fmt.Println(v.IsValid())
	fmt.Println(v.CanFloat())
	//fmt.Println(v.IsNil())
	fmt.Println(v.IsZero())
	fmt.Println("b---")

	var c map[string]bool
	c = make(map[string]bool)
	v = reflect.ValueOf(c)
	fmt.Println(v.IsValid())
	fmt.Println(v.IsZero())
	fmt.Println(v.IsNil())
	fmt.Println("c---")

	v = reflect.ValueOf(nil)
	fmt.Println(v.IsValid())
}

func SetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
