package reflect_struct

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Name   string `json:"name2"`
	Age    int    `json:"age"`
	Gender int8   `json:"gender" db:"gender_db"`
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(n string) {
	u.Name = n
}

func TestStruct(t *testing.T) {
	u := User{Name: "tom", Age: 10}
	field, _ := reflect.TypeOf(u).FieldByName("Name")
	fmt.Printf("%#v\n", field)
	s1 := field.Tag.Get("json")
	fmt.Println(s1, field.Type)
	for i := 0; i < reflect.TypeOf(u).NumField(); i++ {
		tmp := reflect.TypeOf(u).Field(i).Tag.Get("json")
		tmp2 := reflect.TypeOf(u).Field(i).Tag.Get("db")
		fmt.Println(tmp, tmp2)
	}
}

func TestMethod(t *testing.T) {
	u := &User{Name: "tom", Age: 10}
	for i := 0; i < reflect.TypeOf(u).NumMethod(); i++ {
		fmt.Println(reflect.TypeOf(u).Method(i).Name)
	}
}

func TestInvalid(t *testing.T) {
	var u User
	u.Name = "s"
	v := reflect.ValueOf(u)
	for i := 0; i < v.NumField(); i++ {
		fmt.Println(v.Field(i).IsZero())
	}
}
