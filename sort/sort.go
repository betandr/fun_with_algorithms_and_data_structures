package sort

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](items []T) {
	if len(items) < 2 {
		return
	}

	for i := len(items); i > 0; i-- {
		for j := 1; j < i; j++ {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
		}
	}
}

func InsertionSort[T constraints.Ordered](items []T) {
	if len(items) < 2 {
		return
	}

	for i := 1; i < len(items); i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j--
		}
	}
}

func SelectionSort[T constraints.Ordered](items []T) {
	if len(items) < 2 {
		return
	}

	var m int

	for i := 0; i < len(items)-1; i++ {
		m = i

		for j := i + 1; j < len(items); j++ {
			if items[j] < items[m] {
				m = j
			}
		}
		items[m], items[i] = items[i], items[m]
	}

}

func merge[T constraints.Ordered](left []T, right []T) []T {

	var i, j int
	size := len(left) + len(right)
	m := make([]T, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			m[k] = right[j]
			j++

		} else if j > len(right)-1 && i <= len(left)-1 {
			m[k] = left[i]
			i++

		} else if left[i] < right[j] {
			m[k] = left[i]

		} else {
			m[k] = right[j]
			j++
		}
	}

	return m
}

func MergeSort[T constraints.Ordered](items []T) []T {
	if len(items) < 2 {
		return items
	}

	mid := (len(items)) / 2

	return merge(MergeSort(items[:mid]), MergeSort(items[mid:]))
}

func QuickSort[T constraints.Ordered](items []T) {
	quickSort(items, 0, len(items)-1)
}

func quickSort[T constraints.Ordered](items []T, start int, end int) {
	if start < end {
		pi := partition(items, start, end)
		quickSort(items, start, pi-1)
		quickSort(items, pi+1, end)
	}
}

func partition[T constraints.Ordered](items []T, start, end int) int {
	pivot := items[end]
	i := start - 1

	for j := start; j < end; j++ {
		if items[j] <= pivot {
			i++
			items[i], items[j] = items[j], items[i]
		}
	}

	items[i+1], items[end] = items[end], items[i+1]

	return i + 1
}
