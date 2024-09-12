package lesson2

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
	sex  string
}

func NewPerson(name string, age int, sex string) *Person {
	return &Person{
		Name: name,
		Age:  age,
		sex:  sex,
	}
}

func (p *Person) Introduce() string {
	return fmt.Sprintf("我的名字叫%s,今年%d岁, 性别：%s", p.Name, p.Age, p.sex)
}

func (p *Person) SetSex(sex string) {
	p.sex = sex
}

func (p *Person) GetSex() string {
	return p.sex
}

type Teacher struct {
	Person
	Phone   int64
	Address string
}

func NewTeacher(name string, age int, phone int64, address string) *Teacher {
	return &Teacher{
		Person:  *NewPerson(name, age, ""),
		Phone:   phone,
		Address: address,
	}
}

func (t *Teacher) SetTeacher(person *Person, phone int64, address string) {
	t.Person = *person
	t.Phone = phone
	t.Address = address
}

func (t *Teacher) SetTeacher2(name string, phone int64, address string) {
	t.Person.Name = name
	t.Phone = phone
	t.Address = address
}

func (t *Teacher) Introduce2() string {
	return fmt.Sprintf("我的名字叫%s,今年%d岁,我的电话是%d,我的地址是%s", t.Name, t.Age, t.Phone, t.Address)
}

type Product struct {
	Id    int
	Name  string
	Price float64
}

func FindProductByID(products []Product, id int) Product {
	for _, product := range products {
		if product.Id == id {
			return product
		}
	}
	return Product{}
}

type User struct {
	Username string `json:"username"` // 添加 JSON 标签
	Password string `json:"password"`
	Email    string `json:"email"`
}

func TestUser(t *testing.T) {
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

	// 反序列化 JSON 字符串回 User 结构体
	var newUser User
	err = json.Unmarshal(jsonData, &newUser)
	if err != nil {
		fmt.Println("Error unmarshalling from JSON:", err)
		return
	}
}
