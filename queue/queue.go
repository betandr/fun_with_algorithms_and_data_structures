package queue

import (
	"fmt"
	"fun_with_algorithms_and_data_structures/linked_list"
	"strings"
)

type Queue[T any] struct {
	linkedList linked_list.LinkedList[T]
}

func new[T any]() Queue[T] {
	st := Queue[T]{}
	st.linkedList = linked_list.LinkedList[T]{}
	return st
}

func (q *Queue[T]) Enqueue(o T) {
	node := linked_list.Node[T]{}
	node.Value = o
	q.linkedList.AddTail(&node)
}

func (q *Queue[T]) Dequeue() error {
	if q.linkedList.Count == 0 {
		return fmt.Errorf("error: cannot dequeue from an empty queue")
	}
	q.linkedList.RemoveHead()

	return nil
}

func (q *Queue[T]) Peek() (T, error) {
	var r T
	if q.linkedList.Count == 0 {
		return r, fmt.Errorf("error: cannot peek an empty queue")
	}

	r = q.linkedList.Head.Value
	return r, nil
}

func (q *Queue[T]) String() string {
	var sb strings.Builder
	fmt.Fprint(&sb, "START")

	node := q.linkedList.Head
	for node != nil {
		fmt.Fprintf(&sb, " %v", node.Value)
		node = node.Next
	}

	fmt.Fprint(&sb, " END")
	return sb.String()
}
