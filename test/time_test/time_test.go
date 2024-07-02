package time_test

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	tm := time.Now() //获取当前时间
	fmt.Println(tm)
}

func TestLocal(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/London")
	now := time.Now().In(loc)
	fmt.Println(now)
}

func TestFormat(t *testing.T) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("2006+01+02"))
}

func TestParse(t *testing.T) {
	s1 := "2024-05-15 13:11:35"
	t1, err := time.Parse("2006-01-02 15:04:05", s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t1)
	loc, _ := time.LoadLocation("America/New_York")
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", s1, loc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t2)
}

func TestDate(t *testing.T) {
	//2024-05-10 14:30:48
	loc, _ := time.LoadLocation("America/New_York")
	date := time.Date(2024, 5, 10, 14, 30, 48, 0, loc)
	fmt.Println(date)
}

func TestTimestamp(t *testing.T) {
	fmt.Println(time.Now().Unix())      //秒级时间戳
	fmt.Println(time.Now().UnixMilli()) //毫秒级时间戳
	fmt.Println(time.Now().UnixMicro()) //微秒级时间戳
	fmt.Println(time.Now().UnixNano())  //纳秒级时间戳
}

func TestAdd(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Add(time.Hour))           //增加一小时
	fmt.Println(now.Add(time.Hour * 24))      //增加一天
	fmt.Println(now.Add(time.Hour * 24 * -1)) //减少一天
}

func TestSleep(t *testing.T) {

	after := time.After(time.Second * 5)
	for {
		select {
		case <-after:
			fmt.Println("等待5秒结束")
			return
		default:
			time.Sleep(time.Second)
			log.Println("wait")
		}
	}
}
