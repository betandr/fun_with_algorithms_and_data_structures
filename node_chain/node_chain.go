package node_chain

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func (n *Node[value]) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%v", n.Value)

	curr := n
	for curr.Next != nil {
		curr = curr.Next
		fmt.Fprintf(&sb, " %v", curr.Value)
	}

	return sb.String()
}
