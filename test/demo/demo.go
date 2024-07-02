package demo

import "fmt"

var (
	name = "abc"
	Id   = 100
)

func GetName() {
	fmt.Println(name)
}

func setId(n int) {
	Id = n
}

func init() {
	fmt.Println("demo init")
}
