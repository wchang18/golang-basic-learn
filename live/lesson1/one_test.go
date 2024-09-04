package lesson1

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

var userPhone int //可以为全局变量，默认为0

func TestQuestion1(t *testing.T) {

	var userPhone2 = 18622228888 //可以为全局变量，赋值声明
	userPhone3 := 18622228888    //只能用在函数内部
	var (
		userPhone4 int
		userPhone5 int
	)
	t.Log(userPhone, userPhone2, userPhone3, userPhone4, userPhone5)
}

func TestQuestion2(t *testing.T) {
	const (
		Second = 1
		Minute = 60 * Second
		Hour   = 60 * Minute
	)
	//time.Second
	wait10 := 10 * Second    //常量可以自动转换为int
	var wait20 int8 = Minute //常量可以自动转换为int8
	var wait30 int32 = Hour  //常量可以自动转换为int32
	t.Log(wait10, wait20, wait30)
}

func TestQuestion3(t *testing.T) {
	const (
		_ = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		yue
		Saturday
		Sunday
		mon
	)
	t.Log(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday, mon, yue)
}

func TestQuestion4(t *testing.T) {
	i := 1234
	j := 5.43
	s1 := fmt.Sprintf("%v", i)
	s2 := fmt.Sprint(j)

	s3 := strconv.Itoa(i)
	s4 := strconv.FormatInt(int64(i), 16)
	t.Log(s1, s2, s3, s4)
}

func TestQuestion5(t *testing.T) {
	s := "5678"
	n, err := strconv.Atoi(s)
	if err != nil {
		t.Log(err)
	}
	m, _ := strconv.ParseInt(s, 10, 64) //bitSize参数可以0, 8, 16, 32, and 64
	t.Log(n, m)
}

func TestQuestion6(t *testing.T) {
	n := 1<<31 - 1
	t.Log(n, math.MaxInt32)

	a := 1
	for j := 0; j < 31; j++ {
		a *= 2
	}
	a -= 1
	t.Log(a)
}

func TestQuestion7(t *testing.T) {
	s1 := "永远的神啊yyds"
	t.Log(len(s1))
	r1 := []rune(s1)
	t.Log(len(r1))
	for i := len(r1) - 1; i >= 0; i-- {
		fmt.Printf("%c", r1[i])
	}
}

func TestQuestion8(t *testing.T) {
	for i := 1; i < 10; i++ {
		str := ""
		for j := i; j < 10; j++ {
			str += fmt.Sprintf("%d * %d = %2d  ", i, j, i*j)
		}
		fmt.Println(str)
	}
}

func TestFibonacci(t *testing.T) {
	fmt.Println(Fibonacci(20))
	fmt.Println(Fibonacci2(20))
}

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Fibonacci2(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	a, b := 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
