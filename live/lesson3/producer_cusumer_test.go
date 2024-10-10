package lesson3

import (
	"log"
	"testing"
	"time"
)

func Producer(ch chan int, count int) {
	for i := 0; i < count; i++ {
		time.Sleep(time.Millisecond * 50)
		ch <- i
	}
}

func Consumer(ch chan int) {
	for i := range ch {
		log.Println("consumer:", i)
	}
}

func Consumer2(ch, ch2 chan int) {
	for {
		select {
		case i := <-ch:
			log.Println("consumer2-i:", i)
		case j := <-ch2:
			log.Println("consumer2-j:", j)
		}
	}
}

func TestPC(t *testing.T) {
	ch := make(chan int)
	ch2 := make(chan int)
	go Consumer2(ch, ch2)
	go Producer(ch, 100)
	Producer(ch2, 100)
	//time.Sleep(time.Second)
}
