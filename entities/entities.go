package entities

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	Key      string
	Adjacent []*Vertex
}

type Ant struct {
	Id        int
	PathIndex int
	Position  int
	Finished  bool
}
