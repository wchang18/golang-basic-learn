package lesson3

import (
	"log"
	"sync"
	"testing"
	"time"
)

func add(a, b int) int {
	time.Sleep(time.Second)
	log.Println(a, b)
	return a + b
}

func SumEight() int {
	n := 1000
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = i + 1
	}
	log.Println(len(numbers))
	log.Println(addSome(numbers...))
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)
	totalOne := make([]int, 0)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		list := numbers[i*10 : (i+1)*10]
		go func(list []int) {
			defer wg.Done()
			sum := 0
			for _, item := range list {
				sum = add(sum, item)
			}
			mu.Lock()
			defer mu.Unlock()
			totalOne = append(totalOne, sum)
		}(list)
	}
	wg.Wait()
	log.Println("totalOne:", len(totalOne), totalOne)
	log.Println(addSome(totalOne...))
	totalTwo := make([]int, 0)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		list := totalOne[i*10 : (i+1)*10]
		go func(list []int) {
			defer wg.Done()
			sum := 0
			for _, item := range list {
				sum = add(sum, item)
			}
			mu.Lock()
			defer mu.Unlock()
			totalTwo = append(totalTwo, sum)
		}(list)
	}
	wg.Wait()
	log.Println("totalTwo:", len(totalTwo), totalTwo)
	log.Println(addSome(totalTwo...))
	total := 0
	for _, item := range totalTwo {
		total = add(total, item)
	}
	return total
}

func TestEight(t *testing.T) {
	t.Log(SumEight())
}

func addSome(a ...int) int {
	sum := 0
	for _, item := range a {
		sum += item
	}
	return sum
}
