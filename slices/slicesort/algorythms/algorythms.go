package algorythms

import (
	"sort"
)

func BuiltInSort(slice []int) []int {
	sortSlice := make([]int, len(slice))

	_ = copy(sortSlice, slice)

	sort.Ints(sortSlice)

	return sortSlice
}

func SelectionSort(slice []int) []int {
	sortSlice := make([]int, len(slice))

	_ = copy(sortSlice, slice)

	for i := 0; i < len(sortSlice)-1; i++ {
		min := i

		for j := i + 1; j < len(sortSlice); j++ {
			if sortSlice[min] > sortSlice[j] {
				min = j
			}
		}

		sortSlice[i], sortSlice[min] = sortSlice[min], sortSlice[i]
	}

	return sortSlice
}

func BubbleSort(slice []int) []int {
	sortSlice := make([]int, len(slice))

	_ = copy(sortSlice, slice)

	l := len(sortSlice) - 1

	for i := 0; i < l; i++ {
		for j := 0; j < l-i; j++ {
			if sortSlice[j] > sortSlice[j+1] {
				sortSlice[j], sortSlice[j+1] = sortSlice[j+1], sortSlice[j]
			}
		}
	}

	return sortSlice
}

func PancakeSort(slice []int) []int {
	sortSlice := make([]int, len(slice))

	_ = copy(sortSlice, slice)

	flip := func(sl []int, pos int) []int {
		s := make([]int, len(sl))

		_ = copy(s, sl)

		for i := 0; i < pos/2; i++ {
			s[i], s[pos-i] = s[pos-i], s[i]
		}

		return s
	}

	l := len(sortSlice)

	for i := 0; i < len(sortSlice)-1; i++ {
		l--

		max := l

		for j := 0; j < l; j++ {
			if sortSlice[max] < sortSlice[j] {
				max = j
			}
		}

		sortSlice = flip(sortSlice, max)
		sortSlice = flip(sortSlice, l)
	}

	return sortSlice
}

func RecursiveBubbleSort(slice []int, l int) []int {
	if l == 0 {
		return slice
	}

	sortSlice := make([]int, len(slice))

	_ = copy(sortSlice, slice)

	for i := 0; i < l; i++ {
		if sortSlice[i] > sortSlice[i+1] {
			sortSlice[i], sortSlice[i+1] = sortSlice[i+1], sortSlice[i]
		}
	}

	sortSlice = RecursiveBubbleSort(sortSlice, l-1)

	return sortSlice
}
