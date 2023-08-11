package node_chain

import (
	"fmt"
	"testing"
)

type nv int // so we can add a String() func

func (n nv) String() string {
	return fmt.Sprintf("%d", n)
}

func TestNodeChain(t *testing.T) {
	tail := Node[nv]{Value: 3}
	body := Node[nv]{Value: 2, Next: &tail}
	head := Node[nv]{Value: 1, Next: &body}

	want := "1 2 3"
	got := head.String()
	if want != got {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMove(t *testing.T) {

	tail := Node[nv]{Value: 4}
	second := Node[nv]{Value: 3, Next: &tail}
	first := Node[nv]{Value: 2, Next: &second}
	head := Node[nv]{Value: 1, Next: &first}

	first.Next = &tail
	second.Next = &first
	head.Next = &second

	want := "1 3 2 4"
	got := head.String()
	if want != got {
		t.Errorf("got %s, want %s", got, want)
	}
}
