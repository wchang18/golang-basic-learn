package test

import (
	"fmt"
	"testing"
)

func TestIf(t *testing.T) {
	if OK() {
		fmt.Println("true") //如果为true，就执行这里
	} else {
		fmt.Println("false") //如果为false，就执行这里
	}

	if size := getSize(); size > 1000 {
		fmt.Println("大于1000")
	} else if size <= 200 && size >= 101 {
		fmt.Println("小于200大于101")
	} else if size <= 100 && size >= 0 {
		fmt.Println("小于100大于0")
	} else {
		fmt.Println("小于0")
	}
}

func OK() bool {
	return false
}

func getSize() int {
	return 101
}

func TestSwitch(t *testing.T) {
	//checkNum(5)
	//checkNum2(9)
	checkNum3(-1)
}

func checkNum(num int) {
	switch num {
	case 1:
		fmt.Println("等于1")
	case 2:
		fmt.Println("等于2")
	case 3:
		fmt.Println("等于3")
	default:
		fmt.Println("其他值", num)
	}
}

func checkNum2(num int) {
	switch num {
	case 1, 5, 10:
		fmt.Println("等于1,5,10")
	case 2, 6:
		fmt.Println("等于2,6")
	case 3:
		fmt.Println("等于3")
	default:
		fmt.Println("其他值", num)
	}
}

func checkNum3(num int) {
	switch {
	case num >= 1 && num <= 5:
		fmt.Println("1和5之间")
	case num >= 6 && num <= 10:
		fmt.Println("6和10之间")
	case num > 10:
		fmt.Println("大于10")
	default:
		fmt.Println("其他值", num)
	}
}

/*
打印：
其他值 5
等于2,6
大于10
*/
