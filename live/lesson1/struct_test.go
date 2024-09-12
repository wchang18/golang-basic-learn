package lesson1

import (
	"go-basic-learn/live/lesson2"
	"testing"
)

func TestPerson(t *testing.T) {

	p1 := lesson2.NewPerson("张三", 18, "woman")
	t.Log(p1.GetSex())
}
