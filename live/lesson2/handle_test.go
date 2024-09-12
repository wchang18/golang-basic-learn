package lesson2

import (
	"fmt"
	"log"
	"testing"
	"time"
)

type handle func(ask *string, reply *string)

func RecordDuration(origin handle) handle {
	return func(ask *string, reply *string) {
		log.Println("start")
		origin(ask, reply)
		log.Println("end")
	}
}

func TestHandle(t *testing.T) {
	getName := func(ask *string, reply *string) {
		fmt.Println(*ask)
		time.Sleep(time.Second)
		*reply = "my name is gary"
	}
	getName = RecordDuration(getName)
	myAsk := "what is your name"
	youReply := ""
	getName(&myAsk, &youReply)
	t.Log(myAsk)
	t.Log(youReply)
}
