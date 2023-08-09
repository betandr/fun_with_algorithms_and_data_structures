package dijkstra

import "sort"

type PriorityQueue struct {
	keys  []string
	nodes map[string]int
}

func NewQueue() *PriorityQueue {
	var q PriorityQueue
	q.nodes = make(map[string]int)
	return &q
}

func (q *PriorityQueue) Len() int {
	return len(q.keys)
}

func (q *PriorityQueue) Swap(i, j int) {
	q.keys[i], q.keys[j] = q.keys[j], q.keys[i]
}

func (q *PriorityQueue) Less(i, j int) bool {
	a := q.keys[i]
	b := q.keys[j]

	return q.nodes[a] < q.nodes[b]
}

func (q *PriorityQueue) Update(node Node) {
	if _, ok := q.nodes[node.Value]; !ok {
		q.keys = append(q.keys, node.Value)
	}

	q.nodes[node.Value] = node.Distance

	sort.Sort(q)
}

func (q *PriorityQueue) Next() Node {
	val, keys := q.keys[0], q.keys[1:]
	q.keys = keys

	dist := q.nodes[val]
	delete(q.nodes, val)

	return Node{val, dist}
}

func (q *PriorityQueue) Empty() bool {
	return len(q.keys) == 0
}

func (q *PriorityQueue) Peek(val string) (priority int, ok bool) {
	priority, ok = q.nodes[val]
	return
}
