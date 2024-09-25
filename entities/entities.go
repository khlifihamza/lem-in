package entities

type Room struct {
	Name string
}

type Edge struct {
	To   string
	Rev  *Edge
	Cap  int
	Flow int
}

type Graph struct {
	Rooms    map[string]*Room
	Tunnels  map[string][]*Edge
	Start    string
	End      string
	AntCount int
}

type Ant struct {
	Id        int
	PathIndex int
	Position  int
	Finished  bool
}
