package heap

import (
	"testing"
)

func TestExtract(t *testing.T) {

	heap := NewHeap[int]()
	want := 99

	heap.Insert(72)
	heap.Insert(22)
	heap.Insert(want)

	got, ok := heap.Extract()

	if !ok {
		t.Errorf("error extracting item from heap")
	} else {
		if got != 99 {
			t.Errorf("extract: want %d, got %d", want, got)
		}
	}
}
