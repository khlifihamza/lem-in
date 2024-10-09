package entities

// Graph struct holds a list of vertices, where each vertex represents a point in the graph.
type Graph struct {
	Vertices []*Vertex
}

// Vertex struct contains a key (identifier) and a list of adjacent vertices (neighbors).
type Vertex struct {
	Key      string
	Adjacent []*Vertex
}

// Ant struct defines an ant with an ID, its current path index, position in the path, and a flag indicating if it has finished its journey.
type Ant struct {
	Id        int
	PathIndex int
	Position  int
	Finished  bool
}
