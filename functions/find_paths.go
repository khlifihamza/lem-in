package functions

import (
	"container/list"
	"lem-in/entities"
)

func EdmondsKarp(g *Graphh) int {
	flow := 0
	for {
		pred := make(map[string]*entities.Edge)
		q := list.New()
		q.PushBack(g.Start)
		for q.Len() > 0 {
			cur := q.Remove(q.Front()).(string)
			if cur == g.End {
				break
			}
			for _, edge := range g.Tunnels[cur] {
				if pred[edge.To] == nil && edge.To != g.Start && edge.Cap > edge.Flow {
					pred[edge.To] = edge
					q.PushBack(edge.To)
				}
			}
		}
		if pred[g.End] == nil {
			break
		}
		df := 1
		for end := g.End; end != g.Start; {
			edge := pred[end]
			df = min(df, edge.Cap-edge.Flow)
			end = edge.Rev.To
		}
		for end := g.End; end != g.Start; {
			edge := pred[end]
			edge.Flow += df
			edge.Rev.Flow -= df
			end = edge.Rev.To
		}
		flow += df
	}
	return flow
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func FindPaths(g *Graphh, maxPaths int) [][]string {
	paths := [][]string{}
	visited := make(map[string]bool)

	var dfs func(room string, path []string)
	dfs = func(room string, path []string) {
		if len(paths) == maxPaths { // Stop if we've found enough paths
			return
		}
		if room == g.End {
			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			paths = append(paths, pathCopy)
			return
		}

		visited[room] = true
		for _, edge := range g.Tunnels[room] {
			if !visited[edge.To] && edge.Flow > 0 {
				dfs(edge.To, append(path, edge.To))
			}
		}
		visited[room] = false
	}

	dfs(g.Start, []string{g.Start})
	return paths
}
