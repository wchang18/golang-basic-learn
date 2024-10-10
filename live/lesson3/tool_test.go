package lesson3

import (
	"fmt"
	"github.com/wchang18/gotool"
	"testing"
)

func TestArray(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(gotool.InArray(array, 7))
}
