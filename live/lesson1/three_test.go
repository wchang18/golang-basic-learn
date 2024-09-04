package lesson1

import (
	"fmt"
	"log"
	"testing"
)

func TestQuestion21(t *testing.T) {
	t.Log(checkSeven(12345))
	t.Log(checkSeven(123456789))
	t.Log(checkSeven(71122))
	t.Log(checkSeven2(12345))
	t.Log(checkSeven2(123456789))
	t.Log(checkSeven2(71122))
}

func checkSeven(num int) bool {
	size := 10
	str := fmt.Sprint(num)
	for i := 1; i <= len(str); i++ {
		res := num % size
		num /= size
		//println(num)
		if res == 7 {
			return true
		}
	}
	return false
}

func checkSeven2(num int) bool {
	str := fmt.Sprint(num)
	for i := 0; i < len(str); i++ {
		if str[i] == '7' {
			return true
		}
	}
	return false
}

func TestQuestion22(t *testing.T) {
	t.Log(isHuiWen("12321"))
	t.Log(isHuiWen("abc"))
	t.Log(isHuiWen("abcd"))
	t.Log(isHuiWen("abccba"))
}

func isHuiWen(str string) bool {
	half := len(str) / 2
	for i := 0; i < half; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func TestQuestion23(t *testing.T) {
	t.Log(isSushu(17))
	t.Log(isSushu(18))
}

func isSushu(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
		count++
	}
	return true
}

var count int

func TestQuestion24(t *testing.T) {
	list := make([]int, 0)
	for i := 2; i <= 1000; i++ {
		count++
		list = append(list, i)
	}

	n := 0
	for n < len(list) {
		tmp := make([]int, 0)
		for _, item := range list {
			if item%list[n] != 0 {
				tmp = append(tmp, item)
			}
			count++
		}
		list = tmp
		n++
	}
	t.Log(list)
	t.Log(count)
}

//23664

func TestQuestion25(t *testing.T) {
	list := make([]int, 0)
	for i := 2; i <= 1000; i++ {
		if isSushu(i) {
			list = append(list, i)
		}
	}
	t.Log(list)
	t.Log(count)
}

//77191

func GetLogger(level string) func(string) {
	mark := fmt.Sprintf("level:%s", level)
	return func(data string) {
		log.Println(mark, data)
	}
}

func TestMethod(t *testing.T) {
	InfoLogger := GetLogger("info")
	InfoLogger("cmd start")
	InfoLogger("cmd end")
	ErrorLogger := GetLogger("error")
	ErrorLogger("cmd start")
	ErrorLogger("cmd end")
}
