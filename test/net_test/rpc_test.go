package main

import (
	"fmt"
	"net/rpc"
	"testing"
)

func TestClient(t *testing.T) {
	client, err := rpc.Dial("tcp", "localhost:8011")
	if err != nil {
		panic(err)
	}

	var reply string
	err = client.Call("UserService.Say", "tom", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
