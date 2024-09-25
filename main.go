package main

import (
	"fmt"
	"lem-in/functions"
)

func main() {
	g, lines, err := functions.ParseInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	maxFlow := functions.EdmondsKarp(g)
	paths := functions.FindPaths(g, maxFlow)
	if len(paths) == 0 {
		fmt.Println("ERROR: no valid path from start to end")
		return
	}
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println()
	paths = g.Sort(paths)
	functions.SimulateAntMovement(g, paths)
}
