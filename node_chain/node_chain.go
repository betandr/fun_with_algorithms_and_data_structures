package node_chain

import (
	"fmt"
	"strings"
)

type value interface {
	String() string
}

type Node[T value] struct {
	Value T
	Next  *Node[T]
}

func (n *Node[value]) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%s", n.Value)

	curr := n
	for curr.Next != nil {
		curr = curr.Next
		fmt.Fprintf(&sb, " %s", curr.Value)
	}

	return sb.String()
}
