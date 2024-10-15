package functions

import (
	"fmt"

	"lem-in/entities"
)

type Network entities.Graph

// Colony struct represents the ant colony and its main attributes.
// It holds a reference to the graph (network) of rooms and paths, the starting and ending points, and the total number of ants to be deployed.
type Colony struct {
	Graph        *Network
	Start        string
	End          string
	NumberOfAnts int
}

// NewColony creates and returns a new instance of the Colony struct, initializing it with the provided graph,
// start and end points, and the number of ants to be deployed.
func NewColony(Graph *Network, Start, End string, NumberOfAnts int) *Colony {
	return &Colony{
		Graph:        Graph,
		Start:        Start,
		End:          End,
		NumberOfAnts: NumberOfAnts,
	}
}

// AddVertex adds a new vertex with the given key to the Network if it does not already exist,
// if the vertex already exists, it returns an error indicating a duplication issue.
func (g *Network) AddVertex(key string) error {
	if !Contains(g.Vertices, key) {
		g.Vertices = append(g.Vertices, &entities.Vertex{Key: key})
		return nil
	}
	return fmt.Errorf("ERROR: invalid data format, duplicated rooms")
}

// GetVertex retrieves and returns the Vertex associated with the specified key from the Network,
// if no vertex with that key exists, it returns nil.
func (g *Network) GetVertex(key string) *entities.Vertex {
	for i, vertex := range g.Vertices {
		if vertex.Key == key {
			return g.Vertices[i]
		}
	}
	return nil
}

// AddEdge creates a bidirectional connection (tunnel) between two vertices (rooms) in the Network,
// it returns an error if either room does not exist or if the tunnel already exists.
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

// RemoveEdge deletes the bidirectional connection (tunnel) between two vertices (rooms) in the Network.
func (g *Network) RemoveEdge(from, to *entities.Vertex) {
	from.Adjacent = RemoveFromSlice(from.Adjacent, to)
	to.Adjacent = RemoveFromSlice(to.Adjacent, from)
}

// GetShortPath finds the shortest path from the start vertex to the end vertex in the network,
// avoiding the source vertex, and returns the path as a slice of strings.
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

// CheckShortestPaths verifies and modifies the provided shortest paths by removing edges
// between certain rooms and generating new shortest paths, ensuring paths do not return to the source.
func (g *Network) CheckShortestPaths(shortestPaths [][]string, source, end string) [][]string {
	newShortestPaths := [][]string{}
	for i, shortPshortestPath := range shortestPaths {
		if i > 0 {
			for j, room := range shortPshortestPath {
				if j > 0 && ContainsInslice(shortestPaths[0], room) && room != end {
					if len(g.GetVertex(shortPshortestPath[j-1]).Adjacent) > 2 {
						g.RemoveEdge(g.GetVertex(shortPshortestPath[j-1]), g.GetVertex(room))
						newShortestPaths = append(newShortestPaths, g.GetShortPath(shortPshortestPath[0], end, source))
						g.AddEdge(shortPshortestPath[j-1], room)
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

// GetCombination generates all possible paths from the start room to the end room for the given colony,
// avoiding revisiting rooms already present in previous paths and ensuring all paths are unique.
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

// GetPathCombinations creates a map of path combinations for each of the provided shortest paths,
// utilizing the GetCombination method to find all unique paths for each shortest path in the network.
func (g *Network) GetPathCombinations(shortestPaths [][]string, Colony *Colony) map[int][][]string {
	pathCombinations := map[int][][]string{}
	for i, shortshortestPath := range shortestPaths {
		pathCombinations[i] = g.GetCombination(shortshortestPath, Colony)
	}
	return pathCombinations
}
