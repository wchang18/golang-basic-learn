package strconv_test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStrconv(t *testing.T) {

	//数字转字符串
	s1 := strconv.Itoa(123)
	fmt.Printf("%T:%s\n", s1, s1)

	s2 := strconv.FormatInt(189, 10)
	fmt.Printf("%T:%s\n", s2, s2)

	v := 3.1415926535
	s3 := strconv.FormatFloat(v, 'E', -1, 32)
	fmt.Printf("%T:%s\n", s3, s3)
	s4 := strconv.FormatFloat(v, 'E', -1, 64)
	fmt.Printf("%T:%s\n", s4, s4)

	//字符串转数字
	n1, _ := strconv.Atoi("123")
	fmt.Printf("%T:%d\n", n1, n1)

	n11, err := strconv.Atoi("12c")
	fmt.Printf("%T:%d,%v\n", n11, n11, err)

	n2, _ := strconv.ParseInt("456", 10, 32)
	fmt.Printf("%T:%d\n", n2, n2)

	n3, _ := strconv.ParseFloat("3.1415", 32)
	fmt.Printf("%T:%f\n", n3, n3)
}
