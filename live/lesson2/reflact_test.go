package lesson2

import (
	"fmt"
	"reflect"
	"testing"
)

func (u *User) Say(str string) {
	fmt.Println(str, u.Username)
}

func (u *User) GetName() string {
	return u.Username
}

func GetType(val any) string {
	return reflect.TypeOf(val).Name()
}
func GetKind(val any) string {
	return reflect.TypeOf(val).Kind().String()
}

func TestGetType(t *testing.T) {
	a := 1
	b := "hello"
	c := User{}
	d := &User{}
	fmt.Println(GetType(a), GetType(b), GetType(c), GetKind(c), GetKind(d))
}

func GetAllFieldNames(val any) []string {
	list := make([]string, 0)
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return list
	}
	t := reflect.TypeOf(val)
	for i := 0; i < t.Elem().NumField(); i++ {
		list = append(list, t.Elem().Field(i).Name)
	}
	return list
}

func GetAllFieldNames2(val any) []string {
	list := make([]string, 0)
	v := reflect.ValueOf(val)

	if v.Kind() == reflect.Ptr {
		return list
	}

	if v.Kind() != reflect.Struct {
		return list
	}

	t := reflect.TypeOf(val)
	for i := 0; i < t.NumField(); i++ {
		list = append(list, t.Field(i).Name)
	}
	return list
}

func TestGetAllFieldNames(t *testing.T) {
	fmt.Println(GetAllFieldNames(&User{}))
	fmt.Println(GetAllFieldNames2(User{}))
}

func SetStructField(origin any, fieldName string, val any) {
	v := reflect.ValueOf(origin)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return
	}
	v.Elem().FieldByName(fieldName).Set(reflect.ValueOf(val))
}

func TestSetStructField(t *testing.T) {
	u := User{Username: "chang", Password: "12324"}
	SetStructField(&u, "Password", "8888")
	fmt.Println(u)
}

func CallMethod(origin any, methodName string, args ...any) {
	v := reflect.ValueOf(origin)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return
	}
	//fmt.Println(args...)
	values := make([]reflect.Value, 0)
	for _, arg := range args {
		values = append(values, reflect.ValueOf(arg))
	}

	m := v.MethodByName(methodName)
	m.Call(values)
}

func TestCallMethod(t *testing.T) {
	u := User{Username: "chang", Password: "12324"}
	CallMethod(&u, "Say", "hello")
}

func StructToMap(src any) map[string]any {
	var m = make(map[string]any)
	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return m
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		m[field.Name] = v.Field(i).Interface()
	}
	return m
}

func TestStructToMap(t *testing.T) {
	u := User{Username: "chang", Password: "12324"}
	fmt.Println(StructToMap(u))
}
