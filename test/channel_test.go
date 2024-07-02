package test

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	ch := make(chan int, 1) //定义
	ch <- 10                //发送数据
	v, ok := <-ch           //接受数据
	fmt.Println(v, ok)
	close(ch) //关闭通道
	val, ok := <-ch
	fmt.Println(val, ok)
}

func TestNoCache(t *testing.T) {
	ch := make(chan int) // 创建一个无缓存通道

	go func() {
		ch <- 2 // 发送数据到通道
	}()

	value := <-ch // 从通道接收数据
	fmt.Println(value)
}

func TestGetChannelValue(t *testing.T) {
	ch := make(chan int, 5)
	go func() {
		for i := 1; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	getChannelValue(ch)
}

func getChannelValue(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func TestChannel2(t *testing.T) {
	var (
		send    func(int, chan<- int)
		receive func(<-chan int)
	)

	send = func(i int, sendChan chan<- int) {
		sendChan <- i
	}

	receive = func(recvChan <-chan int) {
		res := <-recvChan
		fmt.Println("接收：", res)
	}

	ch := make(chan int)
	go send(11, ch)
	receive(ch)
}
