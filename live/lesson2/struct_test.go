package lesson2

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewPerson(t *testing.T) {
	tom := NewPerson("Tom", 18, "man")
	t.Logf("姓名：%s, 年龄：%d", tom.Name, tom.Age)
	t.Log(tom.sex)
}

func TestIntroduce(t *testing.T) {
	tom := NewPerson("Tom", 18, "woman")
	t.Log(tom.Introduce())
	tom.SetSex("man")
	t.Logf("性别为：%s", tom.GetSex())
}

func TestTeacher(t *testing.T) {
	jim := NewTeacher("Jim", 32, 12345678901, "beijing")
	merry := NewPerson("Merry", 28, "")
	jim.SetTeacher(merry, 12345678902, "shanghai")
	t.Log(jim.Introduce(), jim.Address)
}

func TestTeacher2(t *testing.T) {
	jim := NewTeacher("Jim", 32, 12345678901, "beijing")
	t.Log(jim.Name)
	t.Log(jim.Introduce())
}

func TestFindProductByID(t *testing.T) {
	products := []Product{
		{1, "apple", 5.5},
		{2, "banana", 3.5},
		{3, "orange", 4.5},
	}
	products = append(products, Product{
		Id:   4,
		Name: "mongo",
	})
	t.Logf("%+v", FindProductByID(products, 4))
}

func TestUser2(t *testing.T) {
	// 创建一个 User 实例
	user := User{
		Username: "john_doe",
		Password: "secret_password",
		Email:    "john.doe@example.com",
	}

	// 序列化 User 实例为 JSON 字符串
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println("Serialized JSON:", string(jsonData))

}

func TestUser3(t *testing.T) {
	// 反序列化 JSON 字符串回 User 结构体
	jsonData := `{"username":"john_doe","password":"secret_password","email":"john.doe@example.com"}`
	var newUser User
	err := json.Unmarshal([]byte(jsonData), &newUser)
	if err != nil {
		fmt.Println("Error unmarshalling from JSON:", err)
		return
	}
	t.Logf("%+v", newUser)
}
