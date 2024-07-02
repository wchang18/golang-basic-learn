package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 1010)
		fmt.Println("hello")
	}()
	fmt.Println("你好")
	wg.Wait()
}

func TestGo2(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("执行：", i)
		}(i)
	}
	wg.Wait()
}
