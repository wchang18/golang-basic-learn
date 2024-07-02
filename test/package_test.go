package test

import (
	"fmt"
	dm "go-basic-learn/test/demo"
	"testing"
)

func TestDemo(t *testing.T) {
	dm.GetName()
	fmt.Println(dm.Id)
}

func init() {
	fmt.Println("package test init")
}
