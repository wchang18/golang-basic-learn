package test

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"testing"
)

func TestFloat(t *testing.T) {
	var f1 float32 = 1.1
	f2 := 1.2
	fmt.Println(reflect.TypeOf(f1), reflect.TypeOf(f2), math.MaxFloat32, math.MaxFloat64)
}

func TestComplex(t *testing.T) {
	var (
		c1 complex64
		c2 complex128
	)
	c1 = 1 + 2i
	c2 = 2 + 4i

	c3 := 3 + 6i //complex128

	fmt.Println(c1, c2, c3)

	c := complex(1.2, 3.4) // 构造复数c (1.2+3.4i)
	fc1 := real(c)         //获取复数C的实部 1.2
	fc2 := imag(c)         // 获取复数c的虚部 3.4
	fmt.Println(fc1, fc2)
}

func TestBit(t *testing.T) {
	v1 := 0b00101101 // 代表二进制的 101101，相当于十进制的 45
	v2 := 0o377      // 代表八进制的 377，相当于十进制的 255
	v3 := 0x1p-1     //代表十六进制的 1 除以 2，再除以2，也就是 0.25
	v4 := 0x1a       //十六进制，十进制26
	v5 := 12_45      //数字可用_分割
	fmt.Println(v1, v2, v3, v4, v5)
}

func TestString(t *testing.T) {
	var str string
	str2 := "你好 World"
	str3 := `
			 tom
			 jack
			 martin`

	fmt.Println(str, str2)
	fmt.Println(str3)
}

func TestString2(t *testing.T) {
	s1 := "你好jim"
	s2 := "abcd"
	s3 := s1 + s2                 //两个字符串可以用'+'来链接
	fmt.Println(len(s1), len(s2)) //字符串的长度表示的是字符串的byte数，一个汉字由三个byte组成
	fmt.Println(s3)
	fmt.Println(s2[1], string(s2[1])) //通过字符串下标寻找byte
	fmt.Println(s2 > "abc")           //字符串可以比较大小
	for _, char := range s1 {         //可以通过range来遍历字符串的每个字符，为Unicode字符
		fmt.Println(char, string(char))
	}

	str := "D:\\golang\\learn"
	fmt.Print(str)
}

func TestByte(t *testing.T) {
	var b1 = 'a'      //默认为rune类型
	var b2 byte = 'a' //定义为byte类型
	var b3 rune = '我' //除了ASCII字符，只能定义为rune

	fmt.Println(b1, reflect.TypeOf(b1))
	fmt.Println(b2, reflect.TypeOf(b2))
	fmt.Println(b3, reflect.TypeOf(b3))

	s1 := "你好jim"
	fmt.Println(len([]rune(s1)))
}

func TestConvert(t *testing.T) {
	var a int = 12
	fmt.Printf("%d", a)          //数字转字符串
	fmt.Println(strconv.Itoa(a)) //数字转字符串
	fmt.Println(float64(a))      //数字转浮点类型

	list := []byte{'a', 'b', 'c'}
	fmt.Println(string(list)) //转字符串

	var b int64 = 257
	var c int8 = 34
	fmt.Println(uint8(b), int64(c)) //小类型转大类型，会被截取

	var d = "12"
	fmt.Println(strconv.Atoi(d)) //字符串转数字
}

func TestPoint(t *testing.T) {
	var s string   //定义字符串类型
	var sp *string //字符串的指针类型
	s = "abc"
	fmt.Println(s, sp, &s)
	var i *int
	fmt.Println(i)
	var n int64 = 1
	np := &n
	fmt.Println(reflect.TypeOf(np))

	p := new(int)
	fmt.Println(p)
	fmt.Printf("%T", p)
	fmt.Println("++++++")
	v := 10
	p3 := &v
	fmt.Println(p3)
	fmt.Println(*p3) //*表示对指针取值
	*p3 = 20         //指针赋值
	fmt.Println(v)   //v值变成了20
}
