package lesson4

import "testing"

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

func Sum[T Number | Float](i, j T, k ...T) T {
	total := i + j
	for _, v := range k {
		total += v
	}
	return total
}

func TestSum(t *testing.T) {
	if Sum(1, 2, 3, 4) == 10 {
		t.Log("ok")
	} else {
		t.Error("count error")
		//t.Fatal("sum error")
	}
	t.Log("over")
}

func TestSum2(t *testing.T) {
	t.Run("float", func(t *testing.T) {
		var (
			a = 1.1
			b = 2.2
			c = 3.3
		)
		if Sum(a, b, c) == 6.6 {
			t.Log("ok")
		} else {
			t.Error("error")
		}
	})

	t.Run("int", func(t *testing.T) {
		var (
			a int = 1
			b int = 2
			c int = 3
		)
		if Sum(a, b, c) == 6 {
			t.Log("ok")
		} else {
			t.Error("error")
		}
	})

	t.Run("int8", func(t *testing.T) {
		var (
			a int8 = 1
			b int8 = 2
			c int8 = 3
		)
		if Sum(a, b, c) == 6 {
			t.Log("ok")
		} else {
			t.Error("error")
		}
	})
}
