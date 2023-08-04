package graph

// Node (AKA a Vertex) contains a slice of edges which is all
// of its connections to other nodes.
type Node[T comparable] struct {
	Label T
	Edges []*Edge[T]
}

func NewNode[T comparable](lbl T) Node[T] {
	return Node[T]{Label: lbl}
}

func (n *Node[T]) AddEdge(to *Node[T], weight int) {
	n.Edges = append(n.Edges, &Edge[T]{ToNode: to, Weight: weight})
}

// Edge represents a link between nodes. This is a directed graph
// so an edge just has a node it is connected to.
type Edge[T comparable] struct {
	ToNode  *Node[T]
	Weight  int
	visited bool
}

// unvisited returns all of the edges that haven't been traversed yet
func unvisited[T comparable](edges []*Edge[T]) []*Node[T] {
	unv := []*Node[T]{}
	for _, edge := range edges {
		if !edge.visited {
			unv = append(unv, edge.ToNode)
			edge.visited = true
		}
	}
	return unv
}

// trvse recursively walks through the entire graph
func bfs[T comparable](node *Node[T], labels *[]T) {
	if node != nil {
		*labels = append(*labels, node.Label)
		for _, n := range unvisited(node.Edges) {
			bfs(n, labels)
		}
	}
}

// BreadthFirst traverses through the directed graph visiting nodes
// until it has traveled via all edges.
func BreadthFirst[T comparable](start *Node[T]) []T {
	lbls := []T{}
	bfs(start, &lbls)
	return lbls
}
