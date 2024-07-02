package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrint(t *testing.T) {
	//
	fmt.Print("hello\n")
	fmt.Println("word")
	i := 10
	s1 := fmt.Sprint(i)
	fmt.Println(reflect.TypeOf(s1))
	fmt.Println(true)
}

func TestVar(t *testing.T) {
	type User struct {
		Id   int
		Name string
		Sex  int8
		Del  bool
	}

	user := User{
		Id:   11,
		Name: "tome",
		Sex:  1,
		Del:  false,
	}

	fmt.Println(fmt.Sprintf("%v", user))  //值的默认格式
	fmt.Println(fmt.Sprintf("%+v", user)) //类似%v，但输出结构体时会添加字段名
	fmt.Println(fmt.Sprintf("%#v", user)) //值的Go语法表示
	fmt.Println(fmt.Sprintf("%T", user))  //值的类型的Go语法表示
	fmt.Println(fmt.Sprintf("%T", 1.1))
	fmt.Println(fmt.Sprintf("%t", true)) //布尔值
	fmt.Println(fmt.Sprintf("%T", true))
}

func TestInt(t *testing.T) {
	fmt.Println(fmt.Sprintf("%b", 10))                //表示二进制
	fmt.Println(fmt.Sprintf("%c-%c-%c", 116, 89, 90)) //数字对应的unicode码
	fmt.Println(fmt.Sprintf("%d", 11))                //表示十进制
	fmt.Println(fmt.Sprintf("%o", 15))                //表示八进制
	fmt.Println(fmt.Sprintf("%x", 31))                //表示为十六进制，使用a-f
	fmt.Println(fmt.Sprintf("%X", 31))                //表示为十六进制，使用A-F
	fmt.Println(fmt.Sprintf("%U", 'b'))               //表示为Unicode格式：U+1234
	fmt.Println(fmt.Sprintf("%03d", 5))
}

func TestString(t *testing.T) {
	fmt.Println(fmt.Sprintf("%s", "谁shi"))                //输出字符串
	fmt.Println(fmt.Sprintf("%s", []byte{'a', 'c', 'e'})) //输出字符串
	fmt.Println(fmt.Sprintf("%q", "谁shi\n"))              //
}

func TestFloat(t *testing.T) {
	fmt.Println(fmt.Sprintf("%f", 1235.98765888))    //保留小数点后6位
	fmt.Println(fmt.Sprintf("%F", 1235.98765888))    //保留小数点后6位
	fmt.Println(fmt.Sprintf("%12f", 1235.98765888))  //宽度12，默认精度
	fmt.Println(fmt.Sprintf("%.2f", 1235.98765888))  //默认宽度，精度2
	fmt.Println(fmt.Sprintf("%9.2f", 1235.98765888)) //宽度9，精度2
	fmt.Println(fmt.Sprintf("%9.f", 1235.98765888))  //宽度9，精度0
}
