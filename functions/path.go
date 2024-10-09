package functions

import (
	"fmt"
	"sort"
)

func CleanDuplicatedPaths(shortestPaths [][]string) [][]string {
	for i := 0; i < len(shortestPaths); i++ {
		for j := i + 1; j < len(shortestPaths); j++ {
			if shortestPaths[i][0] == shortestPaths[j][0] && len(shortestPaths[i]) < len(shortestPaths[j]) {
				shortestPaths = DeleteInSlice(shortestPaths, i)
				i--
			}
		}
	}
	return shortestPaths
}

func CleanDuplicatedCombinations(pathCombinations map[int][][]string, Colony *Colony) map[int][][]string {
	for key, pathCombination := range pathCombinations {
		pathCombinations[key] = SortByStartAdjacent(Colony, pathCombination)
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

func calculatePathLimits(paths [][]string, antCount int) []int {
	limits := make([]int, len(paths))
	remainingAnts := antCount
	for {
		pathsUsed := 0
		for i := range paths {
			if remainingAnts == 0 {
				return limits
			}
			if i == 0 || len(paths[i])+limits[i] < len(paths[i-1])+limits[i-1] {
				limits[i]++
				remainingAnts--
				pathsUsed++
			} else {
				break
			}
		}
		if pathsUsed == 0 {
			break
		}
	}

	return limits
}
