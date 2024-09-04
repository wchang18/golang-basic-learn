package lesson1

import (
	"fmt"
	"strings"
	"testing"
)

func TestQuestion11(t *testing.T) {
	//对比数组需要长度和类型一样
	l1 := [...]int32{1, 2, 4, 6}
	l2 := [4]int32{1, 2, 4, 5}
	t.Log(l1 == l2)
	t.Log(l1[len(l1)-2])
}

func TestQuestion12(t *testing.T) {
	var arr [4][4]int64
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			arr[i-1][j-1] = int64(i * j)
		}
	}
	t.Log(arr)
	t.Log(arr[len(arr)-1])
}

func TestQuestion13(t *testing.T) {
	var arr []int
	for i := 1; i <= 5; i++ {
		arr = append(arr, i)
	}
	t.Log(arr)
	arr2 := make([]int, 0)
	for i := 1; i <= 5; i++ {
		arr2 = append(arr2, i)
	}
	t.Log(arr2)
	arr3 := make([]int, 5)
	for i := 1; i <= 5; i++ {
		arr3[i-1] = i
	}
	t.Log(arr3)
}

func TestQuestion14(t *testing.T) {
	var arr []int
	for i := 1; i <= 10; i++ {
		arr = append(arr, i)
	}
	t.Log(arr)
	//冒泡排序
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	t.Log(arr)
}

func sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}

func TestQuestion15(t *testing.T) {
	m1 := make(map[int]int)
	for i := 1; i <= 10; i++ {
		m1[i] = i
	}
	t.Log(m1)
	list := make([]int, 0)
	for i, v := range m1 {
		list = append(list, i)
		t.Log(i, v)
	}
	t.Log(list)
	sort(list)
	for _, n := range list {
		fmt.Print(m1[n])
	}
}

func TestQuestion16(t *testing.T) {
	//将小写字母转成大写字母
	str := "abandon"
	m := map[byte]byte{'a': 'A', 'b': 'B', 'c': 'C', 'd': 'D', 'e': 'E', 'f': 'F', 'g': 'G', 'h': 'H',
		'i': 'I', 'j': 'J', 'k': 'K', 'l': 'L', 'm': 'M', 'n': 'N', 'o': 'O', 'p': 'P', 'q': 'Q',
		'r': 'R', 's': 'S', 't': 'T', 'u': 'U', 'v': 'V', 'w': 'W', 'x': 'X', 'y': 'Y', 'z': 'Z'}
	var newStr []byte
	for i := 0; i < len(str); i++ {
		newStr = append(newStr, m[str[i]])
	}
	t.Log(string(newStr))
	t.Log(strings.ToUpper(str))

	var newStr2 []byte
	for i := 0; i < len(str); i++ {
		newStr2 = append(newStr2, str[i]-32)
	}
	t.Log(string(newStr2))
}

func TestQuestion17(t *testing.T) {
	s1 := "goole"
	s2 := "say hello james"
	t.Log(Repeat(s1))
	t.Log(Repeat(s2))
}

func Repeat(str string) string {
	m := make(map[byte]struct{})
	b := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		if _, ok := m[str[i]]; !ok {
			continue
		}
		b = append(b, str[i])
		m[str[i]] = struct{}{}
	}
	return string(b)
}

func insertNumber(list []int, n int) []int {
	for i := 0; i < len(list); i++ {
		if list[i] > n {
			list = append(list[:i], append([]int{n}, list[i:]...)...)
			return list
		}
	}
	list = append(list, n)
	return list
}

func TestInsert(t *testing.T) {
	list := []int{1, 3, 4, 6, 8, 11, 15}
	t.Log(insertNumber(list, 17))
	t.Log(insertNumber(list, 5))
}
