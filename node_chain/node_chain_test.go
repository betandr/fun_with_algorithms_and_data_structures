package node_chain

import (
	"fmt"
	"testing"
)

type nv int

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
