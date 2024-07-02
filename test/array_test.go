package test

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var arr1 [3]int
	arr2 := [3]uint{2, 3, 4}
	arr3 := [...]string{"aa", "bb"}

	fmt.Println(arr1, arr2, arr3)
}

func TestArray2(t *testing.T) {
	arr := [5]int{1, 2, 3, 4}
	fmt.Println(arr[1], arr[4]) //通过数组索引获取对应的值，默认值为0
	fmt.Println(len(arr))       //获取数组长度
	arr[2] = 10                 //修改数组中的值
	fmt.Println(arr)
	////数组遍历
	for i, item := range arr {
		fmt.Println(i, item)
	}
}

func TestArray3(t *testing.T) {

	two := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(two[1][1], two[2][0])
	two[1][2] = 10
	fmt.Println(two)
}
