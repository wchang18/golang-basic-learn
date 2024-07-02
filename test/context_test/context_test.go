package context_test

import (
	"context"
	"log"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestValueContext(t *testing.T) {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "name", "tom")
	ctx = context.WithValue(ctx, "id", 12)

	t.Log(ctx.Value("name"))
	t.Log(ctx.Value("id"))
}

func TestTimeContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			t.Log("done")
			return
		default:
			time.Sleep(time.Second)
			t.Log("running")
		}
	}
}

func TestDeadlineContext(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			t.Log("done")
			return
		default:
			time.Sleep(time.Second)
			t.Log("running")
		}
	}
}

func TestCancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			Order(ctx, strconv.Itoa(n))
		}(i)
	}

	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
}

func Order(ctx context.Context, id string) {
	for {
		select {
		case <-ctx.Done():
			log.Println("done" + id)
			return
		default:
			time.Sleep(time.Second)
			log.Println("order" + id)
		}
	}
}
