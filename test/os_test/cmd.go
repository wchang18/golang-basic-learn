package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("%#v\n", args)
	if len(args) < 2 {
		return
	}
	switch args[1] {
	case "start":
		fmt.Println("start execute")
	case "stop":
		fmt.Println("stop execute")
	case "restart":
		fmt.Println("restart execute")
	default:
		fmt.Println("not find")
	}
}
