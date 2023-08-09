package heap

type Heap[T any] struct {
	Data []T
}

func NewHeap[T any]() *Heap[T] {
	return &Heap[T]{
		Data: make([]T, 0),
	}
}

func (h *Heap[T]) Push(f func(T, T) bool, v T) {
	h.Data = append(h.Data, v)
	h.heapifyUp(f, (len(h.Data) - 1))
}

func (h *Heap[T]) Pop(f func(T, T) bool) (T, bool) {
	var val T
	if h.Size() == 0 {
		return val, false
	}
	val = h.Data[0]
	h.swap(0, h.Size()-1)
	h.Data = h.Data[:h.Size()-1]
	h.heapifyDown(f, 0)

	return val, true
}

func (h *Heap[T]) Size() int {
	return len(h.Data)
}

func (h *Heap[T]) heapifyUp(f func(T, T) bool, i int) {

	for f(h.Data[parentIndex(i)], h.Data[i]) {
		h.swap(i, parentIndex(i))
		i = parentIndex(i)
	}
}

func (h *Heap[T]) heapifyDown(f func(T, T) bool, i int) {
	l, r, largest := leftChildIndex(i), rightChildIndex(i), i

	if l < len(h.Data) && f(h.Data[i], h.Data[l]) {
		largest = l
	}

	if r < len(h.Data) && f(h.Data[largest], h.Data[r]) {
		largest = r
	}

	if largest != i {
		h.swap(i, largest)
		h.heapifyDown(f, largest)
	}
}

func (h *Heap[T]) swap(i, j int) {
	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
}

func leftChildIndex(i int) int {
	return 2*i + 1
}

func rightChildIndex(i int) int {
	return 2*i + 2
}

func parentIndex(i int) int {
	return (i - 1) / 2
}
