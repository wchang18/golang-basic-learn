package lesson3

import (
	"sync"
	"testing"
	"time"
)

var (
	mux sync.Mutex
)

func Counter(count *int) {
	mux.Lock()
	defer mux.Unlock()
	*count++
}

func TestCounter(t *testing.T) {
	var count int
	for i := 0; i < 1000; i++ {
		go Counter(&count)
	}
	time.Sleep(time.Second)
	t.Log(count)
}

func TimeoutExit(long time.Duration) {
	ticker := time.NewTicker(long)
	//time.After(long)
	select {
	case <-ticker.C:
		panic("timeout exit")
	}
}

func TestTimeoutExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()
	TimeoutExit(time.Second * 3)
}
