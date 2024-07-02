package learn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAds(t *testing.T) {
	a := -5
	if Ads(a) != 5 {
		t.Fatal("Ads计算错误")
	} else {
		t.Log("Ads计算正确")
	}

	if Ads(4) == 4 {
		t.Log("Ads计算正确")
	}
}

func TestSum(t *testing.T) {
	if Sum(3, 5) != 8 {
		t.Fatal("Sum计算错误")
	}
}

func TestMul(t *testing.T) {

	list := []struct {
		Name string
		Num  int
		Res  int
	}{
		{"n1", -3, 3},
		{"n2", -1, 1},
		{"n3", 4, 4},
		{"n4", 8, 8},
	}

	for _, item := range list {
		t.Run(item.Name, func(t *testing.T) {
			if Ads(item.Num) != item.Res {
				t.Fatal(item.Name + "=>error")
			}
		})
	}
}

func TestAs2(t *testing.T) {
	a := 1
	b := 1
	var s *string
	assert.Equal(t, a, b) //判断a,b两个值是否相等
	assert.Nil(t, s)      //判断s是否为nil
	assert.True(t, false) //判断是否为true
}
