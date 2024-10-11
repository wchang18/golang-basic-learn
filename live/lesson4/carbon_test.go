package lesson4

import (
	"github.com/uniplaces/carbon"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	t.Log(now)
	t.Log(now.Add(time.Hour * 24))
	s := now.Format("2006-01-02")
	s += " 00:00:00"
	res, _ := time.Parse("2006-01-02 15:04:05", s)
	t.Log(res)
	res2, _ := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	t.Log(res2)
}

func TestCarbon(t *testing.T) {
	now := carbon.Now()
	t.Log(now)
	t.Log(now.AddDay(), now.AddDays(5), now.AddHours(2))
	t.Log(now.StartOfDay(), now.EndOfDay())
}
