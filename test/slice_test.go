package test

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	s1 := []string{"a", "c", "d", "e", "f", "h"}
	s2 := s1[1:3]
	s3 := s1[3:5]
	fmt.Println(s2, s3)
	s1[1] = "xx" //修改了c
	s1[4] = "yy" //修改了f
	fmt.Println(s2, s3)
}

func TestCreateSlice(t *testing.T) {
	//直接定义
	var s1 []string
	fmt.Printf("%T\n", s1)
	fmt.Println(len(s1), cap(s1))

	//直接赋值
	s2 := []int{1, 2, 3}
	fmt.Println(len(s2), cap(s2))
	//
	////截取方式
	s := []int{1, 2, 3, 4, 5, 6, 7}
	s3 := s[1:4]
	s4 := s[:3]
	s5 := s[5:]
	s6 := s[:]
	fmt.Println(s3, s4, s5, s6)
	//
	////make创建
	s7 := make([]int, 2, 3)
	fmt.Println(s7, len(s7), cap(s7))

}

func TestNilSlice(t *testing.T) {
	var s1 []int
	s2 := []int{}
	s3 := make([]int, 0)

	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	fmt.Println(len(s3), cap(s3))
	fmt.Println(s1 == nil, s2 == nil, s3 == nil)

}

func TestAppend(t *testing.T) {
	var s []int
	s = append(s, 1)
	s = append(s, 2, 3, 4)
	fmt.Println(s)
	s1 := []int{5, 6}
	s = append(s, s1...)
	fmt.Println(s)
}

func TestCopy(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := make([]int, 2, 4)
	copy(s2, s1)
	fmt.Println(s2)
}

func TestDel(t *testing.T) {
	s := []int{1, 2, 3}
	s = append(s[:1], s[2:]...)
	fmt.Println(s)
}
