package binary_tree

import (
	"testing"
)

func TestContains(t *testing.T) {
	tree := new[int]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	want := 3

	if !tree.Contains(want) {
		t.Errorf("%d not found in tree", want)
	}
}

func TestRemoveLeftReplacesNodeParentNil(t *testing.T) {
	tree := new[int]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	want := 1

	if !tree.Contains(want) {
		t.Errorf("%d not found in tree", want)
	}

	tree.Remove(want)

	if tree.Contains(want) {
		t.Errorf("%d not removed from tree", want)
	}
}

func TestRemoveLeftReplacesNode(t *testing.T) {
	tree := new[int]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	want := 2

	if !tree.Contains(want) {
		t.Errorf("%d not found in tree", want)
	}

	tree.Remove(want)

	if tree.Contains(want) {
		t.Errorf("%d not removed from tree", want)
	}
}

func TestRemoveRightChildReplace(t *testing.T) {
	tree := new[int]()

	tree.Add(3)
	tree.Add(2)
	tree.Add(1)

	want := 2

	if !tree.Contains(want) {
		t.Errorf("%d not found in tree", want)
	}

	tree.Remove(want)

	if tree.Contains(want) {
		t.Errorf("%d not removed from tree", want)
	}
}

func TestRemoveRightChildLeftmostChildReplace(t *testing.T) {
	tree := new[int]()

	tree.Add(3)
	tree.Add(6)
	tree.Add(8)
	tree.Add(7)
	tree.Add(9)

	want := 6

	if !tree.Contains(want) {
		t.Errorf("%d not found in tree", want)
	}

	tree.Remove(want)

	if tree.Contains(want) {
		t.Errorf("%d not removed from tree", want)
	}
}

func TestUnbalancedValuesPostOrder(t *testing.T) {
	tree := new[int]()

	want := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)
	tree.Add(4)
	tree.Add(5)
	tree.Add(6)
	tree.Add(7)
	tree.Add(8)
	tree.Add(9)

	got := tree.ValuesPostOrder(tree.Root)

	if got == nil {
		t.Error("returned nil values")
	}

	for i, v := range got {
		if want[i] != got[i] {
			t.Errorf("value at %d: want %d, got %d ", i, want[i], v)
		}
	}
}

func TestBalancedValuesPostOrder(t *testing.T) {
	tree := new[int]()

	want := []int{1, 4, 7, 6, 3, 13, 14, 10, 8}

	tree.Add(8)
	tree.Add(3)
	tree.Add(10)
	tree.Add(1)
	tree.Add(6)
	tree.Add(14)
	tree.Add(4)
	tree.Add(7)
	tree.Add(13)

	got := tree.ValuesPostOrder(tree.Root)

	if got == nil {
		t.Error("returned nil values")
	}

	for i, v := range got {
		if want[i] != got[i] {
			t.Errorf("value at %d: want %d, got %d ", i, want[i], v)
		}
	}
}

func TestUnbalancedValuesPreOrder(t *testing.T) {
	tree := new[int]()

	want := []int{1, 9, 8, 7, 6, 5, 4, 3, 2}

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)
	tree.Add(4)
	tree.Add(5)
	tree.Add(6)
	tree.Add(7)
	tree.Add(8)
	tree.Add(9)

	got := tree.ValuesPreOrder(tree.Root)

	if got == nil {
		t.Error("returned nil values")
	}

	for i, v := range got {
		if want[i] != got[i] {
			t.Errorf("value at %d: want %d, got %d ", i, want[i], v)
		}
	}
}

func TestBalancedValuesPreOrder(t *testing.T) {
	tree := new[int]()

	want := []int{8, 1, 4, 7, 6, 3, 13, 14, 10}

	tree.Add(8)
	tree.Add(3)
	tree.Add(10)
	tree.Add(1)
	tree.Add(6)
	tree.Add(14)
	tree.Add(4)
	tree.Add(7)
	tree.Add(13)

	got := tree.ValuesPreOrder(tree.Root)

	if got == nil {
		t.Error("returned nil values")
	}

	for i, v := range got {
		if want[i] != got[i] {
			t.Errorf("value at %d: want %d, got %d ", i, want[i], v)
		}
	}
}

func TestUnbalancedValuesInOrder(t *testing.T) {
	tree := new[int]()

	want := []int{1, 9, 8, 7, 6, 5, 4, 3, 2}

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)
	tree.Add(4)
	tree.Add(5)
	tree.Add(6)
	tree.Add(7)
	tree.Add(8)
	tree.Add(9)

	got := tree.ValuesInOrder(tree.Root)

	if got == nil {
		t.Error("returned nil values")
	}

	for i, v := range got {
		if want[i] != got[i] {
			t.Errorf("value at %d: want %d, got %d ", i, want[i], v)
		}
	}
}

func TestBalancedValuesInOrder(t *testing.T) {
	tree := new[int]()

	want := []int{1, 4, 7, 6, 3, 8, 13, 14, 10, 2}

	tree.Add(8)
	tree.Add(3)
	tree.Add(10)
	tree.Add(1)
	tree.Add(6)
	tree.Add(14)
	tree.Add(4)
	tree.Add(7)
	tree.Add(13)

	got := tree.ValuesInOrder(tree.Root)

	for i, v := range got {
		if want[i] != got[i] {
			t.Errorf("value at %d: want %d, got %d ", i, want[i], v)
		}
	}
}

func TestRemoveNil(t *testing.T) {
	tree := new[int]()
	err := tree.Remove(3)

	if err == nil {
		t.Error("no error from remove empty item")
	}
}
