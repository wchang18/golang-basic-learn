package lesson4

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/thoas/go-funk"
	"testing"
)

func TestFunk(t *testing.T) {
	//randStr := funk.RandomString(10, []rune{'1', '2', '3', '4', '5', '6', '7', '8',
	//	'9', '0', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'})
	//t.Log(randStr)
	assert.Equal(t, funk.IsEmpty(""), true)
	assert.Equal(t, funk.IsEmpty(nil), true)
}

func TestWithout(t *testing.T) {
	testCases := []struct {
		Arr    interface{}
		Values []interface{}
		Expect interface{}
	}{
		{[]string{"foo", "bar"}, []interface{}{"bar"}, []string{"foo"}},
		{[]int{0, 1, 2, 3, 4}, []interface{}{3, 4}, []int{0, 1, 2}},
		//{[]*funk.Foo{f, b}, []interface{}{b, c}, []*Foo{f}},
	}

	for idx, tt := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			actual := funk.Without(tt.Arr, tt.Values...)
			is.Equal(tt.Expect, actual)
		})
	}
}
