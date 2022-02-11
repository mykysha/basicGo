package slicesort

import (
	"github.com/nndergunov/basicGo/slices/slicesort/algorythms"
	"math/rand"
	"testing"
)

const sliceLength = 10000

func BenchmarkBuiltIn(b *testing.B) {
	slice := createSlice(sliceLength)

	b.ResetTimer()

	_ = algorythms.BuiltInSort(slice)
}

func BenchmarkSelection(b *testing.B) {
	slice := createSlice(sliceLength)

	b.ResetTimer()

	_ = algorythms.SelectionSort(slice)
}

func BenchmarkBubble(b *testing.B) {
	slice := createSlice(sliceLength)

	b.ResetTimer()

	_ = algorythms.BubbleSort(slice)
}

func BenchmarkPancake(b *testing.B) {
	slice := createSlice(sliceLength)

	b.ResetTimer()

	_ = algorythms.PancakeSort(slice)
}

func BenchmarkRecursiveBubble(b *testing.B) {
	slice := createSlice(sliceLength)

	l := len(slice) - 1

	b.ResetTimer()

	_ = algorythms.RecursiveBubbleSort(slice, l)
}

func BenchmarkIterative(b *testing.B) {
	slice := createSlice(sliceLength)

	b.ResetTimer()

	_ = algorythms.IterativeSort(slice)
}

func createSlice(length int) []int {
	var slice []int

	used := make(map[int]bool)

	for i := 0; i < length; i++ {
		var n int

		for {
			n = rand.Intn(length)

			if !used[n] {
				break
			}
		}

		used[n] = true

		slice = append(slice, n)
	}

	return slice
}
