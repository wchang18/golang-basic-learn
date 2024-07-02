package test

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "hello"
	}()

	for {
		select {
		case data := <-ch1:
			fmt.Println("ch1:", data)
		case data := <-ch2:
			fmt.Println("ch2:", data)
		default:
			fmt.Println("default")
			time.Sleep(time.Second)
		}
	}

}

func TestSelectTimeout(t *testing.T) {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 2
	}()

	timeout := time.After(3 * time.Second) // 设置超时时间

	select {
	case data := <-ch:
		fmt.Println("Received:", data)
	case <-timeout:
		fmt.Println("Timeout")
	}
}
