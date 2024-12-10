package functions

import (
	"fmt"
	"slices"
	"sort"
)

// CleanDuplicatedCombinations filters out duplicate path combinations from the pathCombinations map,
// keeping only unique combinations based on their string representation after sorting them by starting vertex adjacency.
func CleanDuplicatedCombinations(pathCombinations map[int][][]string, Colony *Colony) map[int][][]string {
	for key, pathCombination := range pathCombinations {
		pathCombinations[key] = SortByStartAdjacent(Colony, pathCombination)
		pathCombinations[key] = slices.CompactFunc(pathCombination, slices.Equal)
	}
	uniqueCombinations := make(map[string][][]string)
	for _, combination := range pathCombinations {
		combinationStr := fmt.Sprintf("%v", combination)
		uniqueCombinations[combinationStr] = combination
	}
	newPathCombinations := make(map[int][][]string)
	i := 0
	for _, combination := range uniqueCombinations {
		newPathCombinations[i] = combination
		i++
	}
	return newPathCombinations
}

// SortByStartAdjacent sorts the given paths based on the adjacency of their starting vertex to the start of the colony,
// prioritizing paths that lead to adjacent rooms and maintaining stable sorting for paths of equal length.
func SortByStartAdjacent(Colony *Colony, paths [][]string) [][]string {
	adjacentOrder := make(map[string]int)
	for i, neighbor := range Colony.Graph.GetVertex(Colony.Start).Adjacent {
		adjacentOrder[neighbor.Key] = i
	}
	sort.SliceStable(paths, func(i, j int) bool {
		if len(paths[i]) != len(paths[j]) {
			return i < j
		}
		orderI, existsI := adjacentOrder[paths[i][0]]
		orderJ, existsJ := adjacentOrder[paths[j][0]]
		if existsI && existsJ {
			return orderI < orderJ
		}
		if existsI {
			return true
		}
		if existsJ {
			return false
		}
		return i < j
	})
	return paths
}

// calculatePathLimits determines how many ants can be allocated to each path based on their lengths and the total ant count,
// ensuring that shorter paths receive more ants first until all ants are allocated or no more paths can be filled.
func calculatePathLimits(paths [][]string, antCount int) []int {
	limits := make([]int, len(paths))
	remainingAnts := antCount
	for {
		for i := range paths {
			if remainingAnts == 0 {
				return limits
			}
			if i == 0 || len(paths[i])+limits[i] < len(paths[i-1])+limits[i-1] {
				limits[i]++
				remainingAnts--
			} else {
				break
			}
		}
	}
}
