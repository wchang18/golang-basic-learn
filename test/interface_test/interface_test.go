package interface_test

import (
	"fmt"
	"testing"
)

// Sayer接口
type Sayer interface {
	Say()
}

type Mover interface {
	Move()
}

// 猫
type Cat struct {
}

func (Cat) Say() {
	fmt.Println("喵喵~")
}

func (Cat) Move() {
	fmt.Println("抓老鼠")
}

// 狗
type Dog struct {
}

func (Dog) Say() {
	fmt.Println("汪汪~")
}

func (*Dog) Move() {
	fmt.Println("摇尾巴")
}

func ListonSay(a Sayer) {
	a.Say()
}

func TestSay(t *testing.T) {
	cat := new(Cat)
	dog := new(Dog)
	ListonSay(cat)
	ListonSay(dog)
}

func TestValue(t *testing.T) {
	var m Mover
	var s Sayer
	s = Dog{}
	m = &Dog{}
	s.Say()
	m.Move()
}

func TestInference(t *testing.T) {
	var a Mover = &Dog{}
	if v, ok := a.(*Dog); ok {
		fmt.Println(v, "类型正确")
	} else {
		fmt.Println(v, "类型不正确")
	}

	if v, ok := a.(*Cat); ok {
		fmt.Println(v, "类型正确")
	} else {
		fmt.Println(v, "类型不正确")
	}

	m := map[string]interface{}{
		"name": "tom",
		"age":  18,
	}
	fmt.Printf("%T-%v\n", m["name"].(string), m["name"])
	fmt.Printf("%T-%v\n", m["age"].(int), m["age"])
}
