package functions

import (
	"lem-in/entities"
	"sort"
)

type Graphh entities.Graph

func (g *Graphh) AddRoom(name string) {
	if _, exists := g.Rooms[name]; !exists {
		g.Rooms[name] = &entities.Room{Name: name}
	}
}

func (g *Graphh) AddTunnel(room1, room2 string) {
	e1 := &entities.Edge{To: room2, Cap: 1}
	e2 := &entities.Edge{To: room1, Cap: 1}
	e1.Rev = e2
	e2.Rev = e1
	g.Tunnels[room1] = append(g.Tunnels[room1], e1)
	g.Tunnels[room2] = append(g.Tunnels[room2], e2)
}

func (g *Graphh) Sort(paths [][]string) [][]string {
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
	return paths
}
