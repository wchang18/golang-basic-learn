package lesson3

import (
	"log"
	"sync"
	"testing"
	"time"
)

type GoPool struct {
	wg      sync.WaitGroup
	Limit   int
	Channel chan struct{}
}

func NewGoPool(limit int) *GoPool {
	return &GoPool{
		Limit:   limit,
		Channel: make(chan struct{}, limit),
	}
}

func (p *GoPool) Add(n int) {
	for i := 0; i < n; i++ {
		p.Channel <- struct{}{}
	}
	p.wg.Add(n)
}

func (p *GoPool) Done() {
	<-p.Channel
	p.wg.Done()
}

func (p *GoPool) Wait() {
	p.wg.Wait()
}

func TestPool(t *testing.T) {
	pool := NewGoPool(10)
	for i := 0; i < 100; i++ {
		pool.Add(1)
		go func(i int) {
			defer pool.Done()
			time.Sleep(time.Second)
			log.Println("i:", i)
		}(i)
	}
	pool.Wait()
}
