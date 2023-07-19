package stack

import (
	"fmt"
	"fun_with_algorithms_and_data_structures/linked_list"
	"strings"
)

type Stack[T any] struct {
	linkedList linked_list.LinkedList[T]
}

func new[T any]() Stack[T] {
	st := Stack[T]{}
	st.linkedList = linked_list.LinkedList[T]{}
	return st
}

func (s *Stack[T]) Push(o T) {
	node := linked_list.Node[T]{}
	node.Value = o
	s.linkedList.AddHead(&node)
}

func (s *Stack[T]) Peek() (T, error) {
	var r T
	if s.linkedList.Count == 0 {
		return r, fmt.Errorf("error: cannot peek an empty stack")
	}
	return s.linkedList.Head.Value, nil
}

func (s *Stack[T]) Pop() (T, error) {
	var r T
	if s.linkedList.Count == 0 {
		return r, fmt.Errorf("error: cannot pop from empty stack")
	}
	r, err := s.Peek()
	if err != nil {
		return r, err
	}
	s.linkedList.RemoveHead()

	return r, nil
}

func (l *Stack[T]) String() string {
	var sb strings.Builder
	fmt.Fprint(&sb, "TOP")

	node := l.linkedList.Head
	for node != nil {
		fmt.Fprintf(&sb, "\n%v", node.Value)
		node = node.Next
	}
	return sb.String()
}
