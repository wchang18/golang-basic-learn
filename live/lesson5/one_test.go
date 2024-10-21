package lesson5

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestFmt(t *testing.T) {
	i := 123
	var s1 string
	s1 = fmt.Sprint(i)
	s1 = fmt.Sprintf("%d", i)
	s1 = fmt.Sprintf("%v", i)
	s1 = strconv.Itoa(i)
	s1 = strconv.FormatInt(int64(i), 10)
	fmt.Println(s1)
	s2 := "456"
	j, err := strconv.Atoi(s2)
	k, err := strconv.ParseInt(s2, 10, 64)
	fmt.Println(j, err, k)
}

func TestSlice(t *testing.T) {
	list := make([]int, 13)
	var (
		length, capacity int
	)
	for i := 0; i <= 512; i++ {
		list = append(list, i)
		length, capacity = len(list), cap(list)
		fmt.Println("len:", length, "cap:", capacity)
	}
	fmt.Println("len:", length, "cap:", capacity)
	fmt.Println("len:", len(list), "cap:", cap(list))
}

func TestMake(t *testing.T) {
	var (
		a = make([]int, 5)
		b = make(map[string]int)
		c = make(chan int)
	)

	var (
		d = new(int)
	)

	fmt.Println(a, b, c, d)

}

func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()
	t.Log("start")
	panic("some error")
	t.Log("end")
}

func TestSelect(t *testing.T) {

	ticker := time.NewTicker(time.Second * 3)
	select {
	case <-ticker.C:
		fmt.Println("timeout")
	}
}
