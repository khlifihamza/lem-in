package main

import (
	"fmt"

	"lem-in/functions"
)

// This function parses a graph, start and end points, number of ants, and handles any errors.
// It prints the initial data and then initializes a Colony structure, retrieves the shortest paths from the graph, sorts them, cleans duplicates, generates path combinations, deploys the ants, and finally prints the movement of the ant army.
func main() {
	graph, text, Start, End, NumberOfAnts, err := functions.Parser()
	if err != nil {
		fmt.Println(err)
		return
	}
	Colony := *functions.NewColony(graph, Start, End, NumberOfAnts)
	shortestPaths := [][]string{}
	errors := []string{}
	for _, vertex := range graph.GetVertex(Start).Adjacent {
		path, err := graph.GetShortPath(vertex.Key, End, Start)
		if err != nil {
			errors = append(errors, err.Error())
			continue
		}
		shortestPaths = append(shortestPaths, path)
	}
	if len(errors) == len(graph.GetVertex(Start).Adjacent) {
		fmt.Println(errors[0])
		return
	}
	for _, line := range text {
		fmt.Println(line)
	}
	fmt.Println()
	shortestPaths = functions.Sort(shortestPaths)
	shortestPaths = graph.CheckShortestPaths(shortestPaths, Start, End)
	pathCombinations := graph.GetPathCombinations(shortestPaths, &Colony)
	pathCombinations = functions.CleanDuplicatedCombinations(pathCombinations, &Colony)
	movements := functions.DeployAntArmy(pathCombinations, &Colony)
	functions.PrintMovements(movements)
}
