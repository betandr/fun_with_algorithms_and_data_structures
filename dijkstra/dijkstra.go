package dijkstra

type Node struct {
	Value    string
	Distance int
}

type Edge map[string]int

type Graph map[string]Edge

func (g Graph) ShortestPath(from, to string) ([]string, int) {
	var path []string
	var distance int

	visited := make(map[string]bool)
	prev := make(map[string]string)

	q := NewQueue()

	q.Update(Node{from, 0})

	for !q.Empty() {
		n := q.Next()

		if n.Value == to {
			distance = n.Distance

			nVal := n.Value
			for nVal != from {
				path = append(path, nVal)
				nVal = prev[nVal]
			}

			break
		}

		visited[n.Value] = true

		for nVal, nDist := range g[n.Value] {
			if visited[nVal] {
				continue
			}

			if _, ok := q.Peek(nVal); !ok {
				prev[nVal] = n.Value
				q.Update(Node{nVal, n.Distance + nDist})
				continue
			}

			routeLen, _ := q.Peek(nVal)
			nodeLen := n.Distance + nDist

			if nodeLen < routeLen {
				prev[nVal] = n.Value
				q.Update(Node{nVal, nodeLen})
			}
		}
	}

	path = append(path, from)
	flip(path)

	return path, distance
}

func flip(path []string) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}
