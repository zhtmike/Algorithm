package main

import (
	"testing"
	"utils/graph"
)

func TestDijkstraAlgo(t *testing.T) {
	g := make(graph.WAdjlist, 4)
	g[0] = []graph.WEdge{{1, 1}, {2, 4}}
	g[1] = []graph.WEdge{{2, 2}, {3, 6}}
	g[2] = []graph.WEdge{{3, 3}}
	result := DijkstraAlgo(g)
	if result[0] != 0 {
		t.Error("Expected 0 got ", result[0])
	}
	if result[1] != 1 {
		t.Error("Expected 1 got ", result[1])
	}
	if result[2] != 3 {
		t.Error("Expected 3 got ", result[2])
	}
	if result[3] != 6 {
		t.Error("Expected 6 got ", result[3])
	}

}
