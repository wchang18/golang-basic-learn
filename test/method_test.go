package test

import (
	"fmt"
	"testing"
)

/*type Student struct {
	Id   int
	name string
}

func (Student) Say() {
	fmt.Printf("say %s\n", "hello")
}

func (s Student) getName() string {
	return s.name
}

func (s *Student) SetName(name string) {
	s.name = name
}

// 构建函数
func NewStudent(id int, name string) Student {
	return Student{
		Id:   id,
		name: name,
	}
}

func TestStudent(t *testing.T) {
	student1 := NewStudent(1, "小明")
	student1.Say()
	fmt.Println(student1.getName())

	student1.SetName("tom")
	fmt.Println(student1.getName())
}*/

type Person struct {
	Name string
	Age  int
}

func (p *Person) Run() {
	fmt.Println(p.Name + "跑步")
}

type Student struct {
	Person
	Sid int
}

func (s *Student) Learn() {
	fmt.Println(s.Name + "学习")
}

func TestRun(t *testing.T) {
	tom := Student{}
	tom.Run()
	tom.Learn()
}
