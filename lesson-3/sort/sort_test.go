package sort

import "testing"

func makeRange(min int64, max int64) []int64 {
	a := make([]int64, max-min+1)
	for i := range a {
		a[i] = max - min + int64(i)
	}
	return a
}

var slice = makeRange(1, 100000)

func BenchmarkSortBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubbleSort(slice)
	}
}

func BenchmarkSortInserts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertsSort(slice)
	}
}
