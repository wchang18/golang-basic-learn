package test

import (
	"fmt"
	"testing"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Del    bool   `json:"del"`
	Parent *User  `json:"parent"`
}

func TestUser(t *testing.T) {
	//第一种方式
	var u User
	u.Id = 100
	u.Name = "james"
	fmt.Printf("%#v\n", u)

	////第二种方式
	u2 := User{
		Id:    101,
		Email: "help@qq.com",
		Del:   true,
		Parent: &User{
			Id: 1011,
		},
	}
	fmt.Printf("%#v\n", u2)

	////第三种方式，不指定字段，但是这种方式必须要填满，而且顺序不能错，一般不使用这种方式
	u3 := User{102, "jack", "99@qq.com", true, nil}
	fmt.Printf("%#v\n", u3)
}

//type Person struct {
//	Name string
//	Age  int
//}
//
//type Student struct {
//	Person //匿名字段
//	Sid    int
//}
//
//func TestStudent(t *testing.T) {
//	jack := Student{
//		Person: Person{
//			Name: "jack",
//			Age:  6,
//		},
//		Sid: 100,
//	}
//
//	fmt.Println(jack.Name)
//	fmt.Println(fmt.Sprintf("%#v", jack))
//}
