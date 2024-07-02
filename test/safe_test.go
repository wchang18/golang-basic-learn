package test

import (
	"fmt"
	"sync"
	"testing"
)

func TestSafe(t *testing.T) {
	var (
		count int
		wg    sync.WaitGroup
	)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func TestSafe2(t *testing.T) {
	var (
		count int
		wg    sync.WaitGroup
		mu    sync.Mutex
	)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

type Connect struct {
	Config string
}

var con *Connect
var once sync.Once

func GetConn() *Connect {
	once.Do(func() {
		con = &Connect{
			Config: "mysql",
		}
		fmt.Println("get connect")
	})
	return con
}

func TestSafeOnce(t *testing.T) {
	var (
		count int
		wg    sync.WaitGroup
	)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
			GetConn()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func TestSafeMap(t *testing.T) {
	var (
		wg sync.WaitGroup
		m  = sync.Map{}
	)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			m.Store(n, fmt.Sprintf("index-%d", n))
			val, _ := m.Load(n)
			fmt.Println(n, val)
		}(i)
	}
	wg.Wait()
}
