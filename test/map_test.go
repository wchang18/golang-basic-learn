package test

import (
	"fmt"
	"sort"
	"testing"
)

func TestMap(t *testing.T) {
	var m1 map[string]int
	fmt.Println(m1)
	m1 = make(map[string]int)
	m1["name"] = 123
	fmt.Println(m1)

	m2 := map[string]string{
		"name":    "tom",
		"address": "shanghai",
	}
	fmt.Println(m2)
}

func TestMap2(t *testing.T) {
	m := map[string]string{
		"phone": "xiaomi",
		"car":   "Tesla",
	}
	fmt.Println(m)
	delete(m, "car")
	fmt.Println(m)

}

func TestMap3(t *testing.T) {
	m := map[string]string{
		"phone": "xiaomi",
		"car":   "Tesla",
		"name":  "tom",
	}

	if val, ok := m["name"]; ok {
		fmt.Println("name", val)
	} else if _, ok = m["car"]; ok {
		fmt.Println("car found")
	} else {
		fmt.Println("not find key")
	}
}

func TestMap4(t *testing.T) {
	m := map[string]string{
		"phone": "xiaomi",
		"car":   "Tesla",
	}

	for key, value := range m {
		fmt.Println(key, value)
	}
}

func TestMap5(t *testing.T) {
	m := map[string]string{
		"d": "dd",
		"c": "cc",
		"b": "bb",
		"a": "aa",
	}

	list := make([]string, 0)

	for k := range m {
		list = append(list, k) //将key添加到切片中
	}
	fmt.Println(list)
	sort.Strings(list) //对切片中的元素排序
	fmt.Println(list)
	for _, key := range list {
		fmt.Println(m[key]) //按顺序打印
	}
}
