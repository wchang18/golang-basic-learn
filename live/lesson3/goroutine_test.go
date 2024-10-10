package lesson3

import (
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestOne1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newNums := make([]int, 0)
	for _, v := range nums {
		time.Sleep(500 * time.Millisecond)
		println(v)
		newNums = append(newNums, v*2)
	}
	t.Log(newNums)
}

func TestOne2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newNums := make([]int, 0)
	var (
		wg sync.WaitGroup
		//mux sync.Mutex
	)
	for _, v := range nums {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//mux.Lock()
			//defer mux.Unlock()
			//time.Sleep(500 * time.Millisecond)
			newNums = append(newNums, i*2)
		}(v)
	}
	wg.Wait()
	t.Log(newNums)
}

func TestTwo(t *testing.T) {
	var (
		lock  sync.Mutex
		wg    sync.WaitGroup
		start = time.Now().Unix()
	)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fetchResource(&lock, i)
		}(i)
	}
	wg.Wait()
	t.Log("spend time:", time.Now().Unix()-start)
}

func fetchResource(lock *sync.Mutex, i int) {
	lock.Lock()
	defer lock.Unlock()
	num := GetRandNumber(3)
	time.Sleep(time.Second * time.Duration(num))
	log.Println("exec:", i, num)
}

func GetRandNumber(n int) int {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := rd.Intn(n) + 1
	return num
}

func TestThree(t *testing.T) {
	var (
		wg sync.WaitGroup
	)

	strs := []string{"apple", "banana", "cherry", "date", "elderberry"}
	for _, str := range strs {
		wg.Add(1)
		go func(str string) {
			defer wg.Done()
			if CheckA(str) {
				log.Println("str:", str)
			}
		}(str)
	}
	wg.Wait()
}

func CheckA(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == 'a' {
			return true
		}
	}
	return false
}
