package main

import (
	"fmt"
	"os/exec"
)

func main() {
	res, err := exec.Command("ls", "-l").Output()
	fmt.Printf("%s\n%v\n", res, err)

	res2, err2 := exec.Command("sh", "-c", "netstat -an | wc -l").Output()
	fmt.Printf("%s\n%v\n", res2, err2)
}
