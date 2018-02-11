package graph

// WEdge is a data structure with the vertex and weight
type WEdge struct {
	Vertex int
	Weight int
}

// WAdjlist is a the weighted adjacent list
type WAdjlist [][]WEdge
