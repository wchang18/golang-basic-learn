package betch

import "testing"

func add(n int) {
	var list []int
	for i := 0; i < n; i++ {
		list = append(list, i)
	}
}

func add2(n int) {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = i
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(20)
	}
}

func BenchmarkAdd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add2(20)
	}
}
