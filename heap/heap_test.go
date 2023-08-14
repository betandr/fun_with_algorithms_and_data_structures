package heap

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func max[T constraints.Ordered](a, b T) bool {
	return a < b
}

func min[T constraints.Ordered](a, b T) bool {
	return a > b
}

func TestMaxHeapExtract(t *testing.T) {

	heap := NewHeap[int]()
	want := 99

	heap.Insert(max[int], 72)
	heap.Insert(max[int], 22)
	heap.Insert(max[int], want)

	got, ok := heap.Extract(max[int])

	if !ok {
		t.Errorf("error extracting item from heap")
	} else {
		if got != want {
			t.Errorf("extract: want %d, got %d", want, got)
		}
	}
}

func TestMinHeapExtract(t *testing.T) {

	heap := NewHeap[int]()
	want := 22

	heap.Insert(min[int], 72)
	heap.Insert(min[int], 99)
	heap.Insert(min[int], want)

	got, ok := heap.Extract(max[int])

	if !ok {
		t.Errorf("error extracting item from heap")
	} else {
		if got != want {
			t.Errorf("extract: want %d, got %d", want, got)
		}
	}
}

func TestMaxHeapPriority(t *testing.T) {

	heap := NewHeap[int]()
	items := []int{72, 22, 12, 7, 89, 44}

	for _, n := range items {
		heap.Insert(max[int], n)
	}

	got := []int{}
	for range items { // use size of items
		item, ok := heap.Extract(max[int])
		if ok {
			got = append(got, item)
		}
	}

	if got == nil {
		t.Error("returned nil values")
	}

	// highest priority is the highest number for this example
	want := []int{89, 72, 44, 22, 12, 7}
	for i, v := range got {
		if want[i] != got[i] {
			t.Errorf("value at %d: want %d, got %d ", i, want[i], v)
		}
	}
}

func TestMinHeapPriority(t *testing.T) {

	heap := NewHeap[int]()
	items := []int{72, 22, 12, 7, 89, 44}

	for _, n := range items {
		heap.Insert(min[int], n)
	}

	got := []int{}
	for range items { // use size of items
		item, ok := heap.Extract(min[int])
		if ok {
			got = append(got, item)
		}
	}

	if got == nil {
		t.Error("returned nil values")
	}

	// lowest priority is the highest number for this example
	want := []int{7, 12, 22, 44, 72, 89}
	for i, v := range got {
		if want[i] != got[i] {
			t.Errorf("value at %d: want %d, got %d ", i, want[i], v)
		}
	}
}
