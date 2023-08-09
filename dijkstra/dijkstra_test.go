package dijkstra

import (
	"testing"
)

func TestShortestPath(t *testing.T) {

	graph := Graph{
		"London":     {"Birmingham": 202, "Bristol": 190, "Salford": 333},
		"Birmingham": {"London": 202, "Bristol": 142},
		"Bristol":    {"Cardiff": 45, "London": 190, "Birmingham": 142},
		"Cardiff":    {"Bristol": 45},
		"Salford":    {"Glasgow": 340, "Birmingham": 151},
		"Glasgow":    {"Salford": 340},
	}

	gotPath, _ := graph.ShortestPath("Cardiff", "Glasgow")
	wantPath := []string{"Cardiff", "Bristol", "London", "Salford", "Glasgow"}

	if len(gotPath) != len(wantPath) {
		t.Errorf("wanted path %v, got %v", wantPath, gotPath)
	}

	for i, v := range gotPath {
		if wantPath[i] != gotPath[i] {
			t.Errorf("value at %d: want %s, got %s ", i, wantPath[i], v)
		}
	}
}

func TestDistance(t *testing.T) {

	graph := Graph{
		"London":     {"Birmingham": 202, "Bristol": 190, "Salford": 333},
		"Birmingham": {"London": 202, "Bristol": 142},
		"Bristol":    {"Cardiff": 45, "London": 190, "Birmingham": 142},
		"Cardiff":    {"Bristol": 45},
		"Salford":    {"Glasgow": 340, "Birmingham": 151},
		"Glasgow":    {"Salford": 340},
	}

	_, gotDistance := graph.ShortestPath("Cardiff", "Glasgow")

	wantDistance := 908
	if gotDistance != wantDistance {
		t.Errorf("wanted distance %d, got %d", wantDistance, gotDistance)
	}

}
