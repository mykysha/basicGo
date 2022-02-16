package slicesort_test

import (
	"testing"

	"github.com/nndergunov/basicGo/slices/slicesort/algorithms"
)

func TestBuiltIn(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName      string
		unsortedSlice []int
		sortedSlice   []int
	}{
		{
			"Built-in Go sort",
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			sliceSrtd := algorithms.BuiltInSort(test.unsortedSlice)

			for i := 0; i < len(test.sortedSlice); i++ {
				if sliceSrtd[i] != test.sortedSlice[i] {
					t.Errorf("wanted %v, got %v", test.sortedSlice, sliceSrtd)
				}
			}
		})
	}
}

func TestSelection(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName      string
		unsortedSlice []int
		sortedSlice   []int
	}{
		{
			"Selection sort",
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			sliceSrtd := algorithms.SelectionSort(test.unsortedSlice)

			for i := 0; i < len(test.sortedSlice); i++ {
				if sliceSrtd[i] != test.sortedSlice[i] {
					t.Errorf("wanted %v, got %v", test.sortedSlice, sliceSrtd)
				}
			}
		})
	}
}

func TestBubble(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName      string
		unsortedSlice []int
		sortedSlice   []int
	}{
		{
			"Bubble sort",
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			sliceSrtd := algorithms.BubbleSort(test.unsortedSlice)

			for i := 0; i < len(test.sortedSlice); i++ {
				if sliceSrtd[i] != test.sortedSlice[i] {
					t.Errorf("wanted %v, got %v", test.sortedSlice, sliceSrtd)
				}
			}
		})
	}
}

func TestPancake(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName      string
		unsortedSlice []int
		sortedSlice   []int
	}{
		{
			"Pancake sort",
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			sliceSrtd := algorithms.PancakeSort(test.unsortedSlice)

			for i := 0; i < len(test.sortedSlice); i++ {
				if sliceSrtd[i] != test.sortedSlice[i] {
					t.Errorf("wanted %v, got %v", test.sortedSlice, sliceSrtd)
				}
			}
		})
	}
}

func TestRecursiveBubble(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName      string
		unsortedSlice []int
		sortedSlice   []int
	}{
		{
			"Recursive bubble sort",
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			l := len(test.unsortedSlice) - 1

			sliceSrtd := algorithms.RecursiveBubbleSort(test.unsortedSlice, l)

			for i := 0; i < len(test.sortedSlice); i++ {
				if sliceSrtd[i] != test.sortedSlice[i] {
					t.Errorf("wanted %v, got %v", test.sortedSlice, sliceSrtd)
				}
			}
		})
	}
}

func TestIterative(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName      string
		unsortedSlice []int
		sortedSlice   []int
	}{
		{
			"Iterative sort",
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			sliceSrtd := algorithms.IterativeSort(test.unsortedSlice)

			for i := 0; i < len(test.sortedSlice); i++ {
				if sliceSrtd[i] != test.sortedSlice[i] {
					t.Errorf("wanted %v, got %v", test.sortedSlice, sliceSrtd)
				}
			}
		})
	}
}
