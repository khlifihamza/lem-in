package functions

import (
	"fmt"

	"lem-in/entities"
)

type Network entities.Graph

type Colony struct {
	Graph        *Network
	Start        string
	End          string
	NumberOfAnts int
}

func NewColony(Graph *Network, Start, End string, NumberOfAnts int) *Colony {
	return &Colony{
		Graph:        Graph,
		Start:        Start,
		End:          End,
		NumberOfAnts: NumberOfAnts,
	}
}

func (g *Network) AddVertex(key string) error {
	if !Contains(g.Vertices, key) {
		g.Vertices = append(g.Vertices, &entities.Vertex{Key: key})
		return nil
	}
	return fmt.Errorf("ERROR: invalid data format, duplicated rooms")
}

func (g *Network) GetVertex(key string) *entities.Vertex {
	for i, vertex := range g.Vertices {
		if vertex.Key == key {
			return g.Vertices[i]
		}
	}
	return nil
}

func (g *Network) AddEdge(from, to string) error {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	if fromVertex == nil || toVertex == nil {
		if fromVertex == nil {
			return fmt.Errorf("ERROR: invalid data format, room %s don't exist", from)
		}
		return fmt.Errorf("ERROR: invalid data format, room %s don't exist", to)
	} else if Contains(fromVertex.Adjacent, to) || Contains(toVertex.Adjacent, from) {
		return fmt.Errorf("ERROR: invalid data format, duplicated tunnels")
	} else {
		fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
		toVertex.Adjacent = append(toVertex.Adjacent, fromVertex)
	}
	return nil
}

func (g *Network) RemoveEdge(from, to *entities.Vertex) {
	from.Adjacent = RemoveFromSlice(from.Adjacent, to)
	to.Adjacent = RemoveFromSlice(to.Adjacent, from)
}

func (g *Network) GetShortPath(start, end, source string) []string {
	queue := [][]string{{start}}
	visited := make(map[string]bool)
	visited[start] = true
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		node := path[len(path)-1]
		if node == end {
			return path
		}
		for _, neighbor := range g.GetVertex(node).Adjacent {
			if neighbor.Key == source {
				continue
			}
			if !visited[neighbor.Key] {
				visited[neighbor.Key] = true
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor.Key)
				queue = append(queue, newPath)
			}
		}
	}
	return []string{}
}

func (g *Network) CheckShortestPaths(shortestPaths [][]string, source, end string) [][]string {
	newShortestPaths := [][]string{}
	for i, shortPshortestPath := range shortestPaths {
		if i > 0 {
			for j, room := range shortPshortestPath {
				if j > 0 && ContainsInslice(shortestPaths[0], room) && room != end {
					if len(g.GetVertex(shortPshortestPath[j-1]).Adjacent) > 2 {
						g.RemoveEdge(g.GetVertex(shortPshortestPath[j-1]), g.GetVertex(room))
						newShortestPaths = append(newShortestPaths, g.GetShortPath(shortPshortestPath[0], end, source))
						break
					}
				}
			}
		}
	}
	for _, newPath := range newShortestPaths {
		if len(newPath) > 0 {
			shortestPaths = append(shortestPaths, newPath)
		}
	}
	shortestPaths = Sort(shortestPaths)
	return shortestPaths
}

func (g *Network) GetCombination(path []string, Colony *Colony) [][]string {
	Combination := [][]string{}
	visited := make(map[string]bool)
	Combination = append(Combination, path)
	for _, vertex := range Colony.Graph.GetVertex(Colony.Start).Adjacent {
		for _, Path := range Combination {
			for _, room := range Path {
				if room != Colony.End {
					visited[room] = true
				}
			}
		}
		if !visited[vertex.Key] {
			queue := [][]string{{vertex.Key}}
			for len(queue) > 0 {
				path := queue[0]
				queue = queue[1:]
				node := path[len(path)-1]
				if node == Colony.End {
					Combination = append(Combination, path)
					visited = make(map[string]bool)
					break
				}
				for _, neighbor := range g.GetVertex(node).Adjacent {
					if neighbor.Key == Colony.Start {
						continue
					}
					if !visited[neighbor.Key] {
						visited[neighbor.Key] = true
						newPath := append([]string{}, path...)
						newPath = append(newPath, neighbor.Key)
						queue = append(queue, newPath)
					}
				}
			}
		}
	}
	Combination = Sort(Combination)
	return Combination
}

func (g *Network) GetPathCombinations(shortestPaths [][]string, Colony *Colony) map[int][][]string {
	pathCombinations := map[int][][]string{}
	for i, shortshortestPath := range shortestPaths {
		pathCombinations[i] = g.GetCombination(shortshortestPath, Colony)
	}
	return pathCombinations
}
