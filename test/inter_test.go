package test

import (
	"fmt"
	"testing"
)

type Mover interface {
	Move()
}

type Cat struct {
}

// 值接收者
func (Cat) Move() {
	fmt.Println("抓老鼠")
}

type Dog struct {
}

// 指针接收者
func (*Dog) Move() {
	fmt.Println("摇尾巴")
}

func SeeMove(m Mover) {
	m.Move()
}

func TestMove(t *testing.T) {
	cat := &Cat{}
	SeeMove(cat)

	dog := &Dog{}
	SeeMove(dog)
}

type anyVal interface{}

type some struct{}

func TestEmpty(t *testing.T) {

	var s anyVal = "hello"
	fmt.Printf("%T\n", s)

	var i anyVal = 10
	fmt.Printf("%T\n", i)

	var v anyVal = some{}
	fmt.Printf("%T\n", v)
}

func TestMapInterface(t *testing.T) {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["name"] = "tom"
	m["age"] = 20
	m["delete"] = false
	fmt.Printf("%#v", m)
}
