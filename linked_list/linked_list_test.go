package linked_list

import "testing"

func TestEmptyLinkedList(t *testing.T) {

	var ll LinkedList[string]

	want := "HEAD->TAIL"
	got := ll.String()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}

func TestAddToHead(t *testing.T) {

	var ll LinkedList[string]

	ll.AddHead(&Node[string]{Value: "ONE"})

	want := "HEAD->ONE->TAIL"
	got := ll.String()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}

func TestAddMultipleToHead(t *testing.T) {

	var ll LinkedList[string]

	ll.AddHead(&Node[string]{Value: "THREE"})
	ll.AddHead(&Node[string]{Value: "TWO"})
	ll.AddHead(&Node[string]{Value: "ONE"})

	want := "HEAD->ONE->TWO->THREE->TAIL"
	got := ll.String()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}

func TestAddToTail(t *testing.T) {
	var ll LinkedList[string]

	ll.AddTail(&Node[string]{Value: "ONE"})

	want := "HEAD->ONE->TAIL"
	got := ll.String()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMultipleAddToTail(t *testing.T) {
	var ll LinkedList[string]

	ll.AddTail(&Node[string]{Value: "ONE"})
	ll.AddTail(&Node[string]{Value: "TWO"})
	ll.AddTail(&Node[string]{Value: "THREE"})

	want := "HEAD->ONE->TWO->THREE->TAIL"
	got := ll.String()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestRemoveHead(t *testing.T) {
	var ll LinkedList[string]

	ll.AddTail(&Node[string]{Value: "ONE"})
	ll.AddTail(&Node[string]{Value: "TWO"})
	ll.AddTail(&Node[string]{Value: "THREE"})

	ll.RemoveHead()

	want := "HEAD->TWO->THREE->TAIL"
	got := ll.String()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestRemoveTail(t *testing.T) {
	var ll LinkedList[string]

	ll.AddTail(&Node[string]{Value: "ONE"})
	ll.AddTail(&Node[string]{Value: "TWO"})
	ll.AddTail(&Node[string]{Value: "THREE"})

	ll.RemoveTail()

	want := "HEAD->ONE->TWO->TAIL"
	got := ll.String()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestEnumerate(t *testing.T) {
	var ll LinkedList[string]

	ll.AddTail(&Node[string]{Value: "1"})
	ll.AddTail(&Node[string]{Value: "2"})
	ll.AddTail(&Node[string]{Value: "3"})

	got := []string{}

	enum := ll.enumerator()

	node := enum.getNext()
	for node != nil {
		got = append(got, node.Value)
		node = enum.getNext()
	}

	if got[0] != "1" || got[1] != "2" || got[2] != "3" {
		t.Errorf("got [%s, %s, %s], want [1, 2, 3]", got[0], got[1], got[2])
	}
}
