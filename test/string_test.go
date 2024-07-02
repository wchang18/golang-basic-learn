package test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestString(t *testing.T) {
	var str string
	str2 := "你好 World"
	str3 := `tom
			 jack
			 martin`

	fmt.Println(str, str2, str3)
}

func TestString2(t *testing.T) {
	s1 := "你好jim"
	s2 := "abcd\n"
	s3 := s1 + s2                 //两个字符串可以用'+'来链接
	fmt.Println(len(s1), len(s2)) //字符串的长度表示的是字符串的byte数，一个汉字由三个byte组成
	fmt.Println(s3)
	fmt.Println(s2[1])        //通过字符串下标寻找byte
	fmt.Println(s2 > "abc")   //字符串可以比较大小
	for _, char := range s1 { //可以通过range来遍历字符串的每个字符，为Unicode字符
		fmt.Println(char, string(char))
	}
	fmt.Print(s2)
}

func TestByte(t *testing.T) {
	var b1 = 'a'      //默认为rune类型
	var b2 byte = 'a' //定义为byte类型
	var b3 rune = '我' //除了ASCII字符，只能定义为rune

	fmt.Println(b1, reflect.TypeOf(b1))
	fmt.Println(b2, reflect.TypeOf(b2))
	fmt.Println(b3, reflect.TypeOf(b3))

	s := "我的abc"
	fmt.Println(len(s), len([]rune(s)))
}

func TestConvert(t *testing.T) {
	var a int = 12
	fmt.Printf("%d", a)          //数字转字符串
	fmt.Println(strconv.Itoa(a)) //数字转字符串
	fmt.Println(float64(a))      //数字转浮点类型

	list := []byte{'a', 'b', 'c'}
	fmt.Println(string(list)) //转字符串
	//
	var b int64 = 257
	var c int8 = 34
	fmt.Println(uint8(b), int64(c)) //小类型转大类型，会被截取
	//
	var d = "12"
	fmt.Println(strconv.Atoi(d)) //字符串转数字
}

/*
97 int32
97 uint8
25105 int32
*/

/*
打印：
9 4
你好jimabcd
98
true
20320 你
22909 好
106 j
105 i
109 m
*/
