package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {

	items := []int{17, 3, 15, 14, 18, 4, 8, 53, 2, 16, 10, 5, 20, 100, 1, 12, 21, 13, 9, 19, 7, 11, 6}
	BubbleSort[int](items)
	lower := 0

	for _, item := range items {
		if item < lower {
			t.Errorf("item %d found before %d", item, lower)
		}
		lower = item
	}
}

func TestInsertionSort(t *testing.T) {
	items := []int{17, 3, 15, 14, 18, 4, 8, 53, 2, 16, 10, 5, 20, 100, 1, 12, 21, 13, 9, 19, 7, 11, 6}
	InsertionSort[int](items)
	lower := 0

	for _, item := range items {
		if item < lower {
			t.Errorf("item %d found before %d", item, lower)
		}
		lower = item
	}
}

func TestSelectionSort(t *testing.T) {
	items := []int{17, 3, 15, 14, 18, 4, 8, 53, 2, 16, 10, 5, 20, 100, 1, 12, 21, 13, 9, 19, 7, 11, 6}
	SelectionSort[int](items)
	lower := 0

	for _, item := range items {
		if item < lower {
			t.Errorf("item %d found before %d", item, lower)
		}
		lower = item
	}
}

func TestMergeSort(t *testing.T) {
	items := []int{17, 3, 15, 14, 18, 4, 8, 53, 2, 16, 10, 5, 20, 100, 1, 12, 21, 13, 9, 19, 7, 11, 6}
	got := MergeSort[int](items)
	lower := 0

	for _, item := range got {
		if item < lower {
			t.Errorf("item %d found before %d", item, lower)
		}
		lower = item
	}
}

func TestQuickSort(t *testing.T) {
	items := []int{17, 3, 15, 14, 18, 4, 8, 53, 2, 16, 10, 5, 20, 100, 1, 12, 21, 13, 9, 19, 7, 11, 6}
	QuickSort[int](items)
	lower := 0

	for _, item := range items {
		if item < lower {
			t.Errorf("item %d found before %d", item, lower)
		}
		lower = item
	}
}
