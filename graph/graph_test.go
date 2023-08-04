package graph

import (
	"testing"
)

func TestTraverseGraph(t *testing.T) {

	nodeA := NewNode("A")
	nodeB := NewNode("B")
	nodeC := NewNode("C")
	nodeD := NewNode("D")
	nodeE := NewNode("E")
	nodeF := NewNode("F")
	nodeG := NewNode("G")

	nodeA.AddEdge(&nodeB, 3)
	nodeA.AddEdge(&nodeC, 3)

	nodeB.AddEdge(&nodeC, 6)
	nodeB.AddEdge(&nodeD, 5)

	nodeC.AddEdge(&nodeD, 11)
	nodeC.AddEdge(&nodeE, 8)

	nodeD.AddEdge(&nodeE, 2)
	nodeD.AddEdge(&nodeF, 10)

	nodeD.AddEdge(&nodeG, 2)

	nodeE.AddEdge(&nodeG, 5)
	nodeF.AddEdge(&nodeG, 3)

	labels := BreadthFirst[string](&nodeA)
	numEdges := 12

	if len(labels) != numEdges {
		t.Errorf("want %d labels from %d edges, got %d", numEdges, numEdges, len(labels))
	}
}
