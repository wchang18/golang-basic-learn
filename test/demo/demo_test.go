package demo

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	GetName()
	setId(300)
	fmt.Println(Id)
}
