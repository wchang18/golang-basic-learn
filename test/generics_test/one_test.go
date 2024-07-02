package generics_test

import (
	"fmt"
	"testing"
)

func InArray[T comparable](v T, s []T) bool {
	for _, item := range s {
		if item == v {
			return true
		}
	}
	return false
}

func TestInArray(t *testing.T) {
	var v float64 = 1.1
	list := []float64{1.2, 1.3, 1.1}
	fmt.Println(InArray[float64](v, list))

	s := "a"
	ss := []string{"aa", "bb", "cc"}
	fmt.Println(InArray[string](s, ss))
}
