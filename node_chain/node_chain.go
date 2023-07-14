package node_chain

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func NodesToString(n *Node[string]) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%s", n.Value)

	curr := n
	for curr.Next != nil {
		curr = curr.Next
		fmt.Fprintf(&sb, " %s", curr.Value)
	}

	return sb.String()
}
