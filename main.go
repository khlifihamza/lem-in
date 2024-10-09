package main

import (
	"fmt"

	"lem-in/functions"
)

func main() {
	graph, text, Start, End, NumberOfAnts, err := functions.Parser()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, line := range text {
		fmt.Println(line)
	}
	fmt.Println()
	Colony := *functions.NewColony(graph, Start, End, NumberOfAnts)
	shortestPaths := [][]string{}
	for _, vertex := range graph.GetVertex(Start).Adjacent {
		shortestPaths = append(shortestPaths, graph.GetShortPath(vertex.Key, End, Start))
	}
	shortestPaths = functions.Sort(shortestPaths)
	shortestPaths = graph.CheckShortestPaths(shortestPaths, Start, End)
	shortestPaths = functions.CleanDuplicatedPaths(shortestPaths)
	pathCombinations := graph.GetPathCombinations(shortestPaths, &Colony)
	pathCombinations = functions.CleanDuplicatedCombinations(pathCombinations, &Colony)
	movements := functions.DeployAntArmy(pathCombinations, &Colony)
	functions.PrintMovements(movements)
}
