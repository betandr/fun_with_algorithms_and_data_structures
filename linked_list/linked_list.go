package linked_list

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Head  *Node[T]
	Tail  *Node[T]
	Count int
}

// addHead adds a node to the head of the list.
func (l *LinkedList[T]) addHead(n *Node[T]) {
	temp := l.Head
	l.Head = n
	l.Head.Next = temp
	l.Count++
	if l.Count == 1 {
		l.Tail = l.Head
	}
}

// addTail adds a node to the tail of the list.
func (l *LinkedList[T]) addTail(n *Node[T]) {
	if l.Count == 0 {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
	}

	l.Tail = n

	l.Count++
}

// removeHead removes a node from the head of the list.
func (l *LinkedList[T]) removeHead() {
	if l.Count != 0 {
		l.Head = l.Head.Next
		l.Count--
		if l.Count == 0 {
			l.Tail = nil
		}
	}

}

// removeTail removes a node from the tail of the list.
// Loops through entire list so decreases in performance by O(n)
func (l *LinkedList[T]) removeTail() {
	if l.Count != 0 {
		if l.Count == 1 {
			l.Head = nil
			l.Tail = nil
		} else {
			current := l.Head
			for current.Next != l.Tail {
				current = current.Next
			}
			current.Next = nil
			l.Tail = current
		}
		l.Count--
	}

}

func (l *LinkedList[T]) String() string {
	var sb strings.Builder
	fmt.Fprint(&sb, "HEAD")

	node := l.Head
	for node != nil {
		fmt.Fprintf(&sb, "->%v", node.Value)
		node = node.Next
	}

	fmt.Fprint(&sb, "->TAIL")
	return sb.String()
}
