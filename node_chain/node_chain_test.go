package node_chain

import (
	"testing"
)

func TestNodeChain(t *testing.T) {
	tail := Node[string]{Value: "World!"}
	body := Node[string]{Value: "there", Next: &tail}
	head := Node[string]{Value: "Hello", Next: &body}

	want := "Hello there World!"
	got := NodesToString(&head)
	if want != got {
		t.Errorf("got %s, want %s", got, want)
	}
}
