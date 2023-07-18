package linked_list

type Enumerator[T any] struct {
	Current *Node[T]
	List    *LinkedList[T]
}

func (e *Enumerator[T]) getNext() *Node[T] {
	if e.Current == nil {
		e.Current = e.List.Head
	} else {
		temp := e.Current.Next
		e.Current = temp
	}
	return e.Current
}

func (l *LinkedList[T]) enumerator() Enumerator[T] {
	return Enumerator[T]{
		List: l,
	}
}
