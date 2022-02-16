package slicesort_test

import (
	"math/rand"
	"testing"

	"github.com/nndergunov/basicGo/slices/slicesort/algorithms"
)

const sliceLength = 10000

func BenchmarkBuiltIn(b *testing.B) {
	slice := createSlice(sliceLength)

	b.ResetTimer()

	_ = algorithms.BuiltInSort(slice)
}

func BenchmarkSelection(b *testing.B) {
	slice := createSlice(sliceLength)

	b.ResetTimer()

	_ = algorithms.SelectionSort(slice)
}

func BenchmarkBubble(b *testing.B) {
	slice := createSlice(sliceLength)

	b.ResetTimer()

	_ = algorithms.BubbleSort(slice)
}

func BenchmarkPancake(b *testing.B) {
	slice := createSlice(sliceLength)

	b.ResetTimer()

	_ = algorithms.PancakeSort(slice)
}

func BenchmarkRecursiveBubble(b *testing.B) {
	slice := createSlice(sliceLength)

	l := len(slice) - 1

	b.ResetTimer()

	_ = algorithms.RecursiveBubbleSort(slice, l)
}

func BenchmarkIterative(b *testing.B) {
	slice := createSlice(sliceLength)

	b.ResetTimer()

	_ = algorithms.IterativeSort(slice)
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
