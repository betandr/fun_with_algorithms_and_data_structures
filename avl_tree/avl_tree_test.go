package avl_tree

import (
	"testing"
)

func TestRandomDataIsInOrder(t *testing.T) {
	tree := NewAVLTree[int]()

	tree.Add(98)
	tree.Add(22)
	tree.Add(14)
	tree.Add(72)
	tree.Add(44)
	tree.Add(25)
	tree.Add(63)

	want := []int{14, 22, 25, 44, 63, 72, 98} // ordered slice
	got := tree.InOrder()

	if len(got) != len(want) {
		t.Errorf("got slice of size %d, want size %d ", len(got), len(want))
	} else {
		for i, v := range got {
			if want[i] != got[i] {
				t.Errorf("value at %d: want %d, got %d ", i, want[i], v)
			}
		}
	}
}
