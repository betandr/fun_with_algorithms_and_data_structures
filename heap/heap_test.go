package heap

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func less[T constraints.Ordered](a, b T) bool {
	return a < b
}

func TestExtract(t *testing.T) {

	heap := NewHeap[int]()
	want := 99

	heap.Push(less[int], 72)
	heap.Push(less[int], 22)
	heap.Push(less[int], want)

	got, ok := heap.Pop(less[int])

	if !ok {
		t.Errorf("error extracting item from heap")
	} else {
		if got != 99 {
			t.Errorf("extract: want %d, got %d", want, got)
		}
	}
}

func TestPriority(t *testing.T) {

	heap := NewHeap[int]()
	items := []int{72, 22, 12, 7, 89, 44}

	for _, n := range items {
		heap.Push(less[int], n)
	}

	got := []int{}
	for range items { // use size of items
		item, ok := heap.Pop(less[int])
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
