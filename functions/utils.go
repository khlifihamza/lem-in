package functions

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"lem-in/entities"
)

func Contains(slice []*entities.Vertex, key string) bool {
	for _, vertex := range slice {
		if key == vertex.Key {
			return true
		}
	}
	return false
}

func RemoveFromSlice(slice []*entities.Vertex, vertex *entities.Vertex) []*entities.Vertex {
	for i, v := range slice {
		if v == vertex {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func ContainsInslice(slice []string, key string) bool {
	for _, vertex := range slice {
		if key == vertex {
			return true
		}
	}
	return false
}

func DeleteInSlice(shortestPaths [][]string, index int) [][]string {
	newSlice := [][]string{}
	for i, slice := range shortestPaths {
		if i != index {
			newSlice = append(newSlice, slice)
		}
	}
	return newSlice
}

func Sort(paths [][]string) [][]string {
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
	return paths
}

func GetNumberOfSteps(result [][]string) int {
	c := 0
	for _, slice := range result {
		c += len(slice)
	}
	return c
}

func PrintMovements(movements [][]string) {
	for _, movement := range movements {
		for i, step := range movement {
			fmt.Print(step)
			if i < len(movement)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func CompareAllResults(allResults map[int][][]string) int {
	index := 0
	minTurns := len(allResults[0])
	minSteps := GetNumberOfSteps(allResults[0])
	for i := 0; i < len(allResults); i++ {
		if len(allResults[i]) == minTurns && len(allResults[i]) > 0 {
			if minSteps > GetNumberOfSteps(allResults[i]) {
				minTurns = len(allResults[i])
				minSteps = GetNumberOfSteps(allResults[i])
				index = i
			}
		} else if len(allResults[i]) < minTurns && len(allResults[i]) > 0 {
			minTurns = len(allResults[i])
			minSteps = GetNumberOfSteps(allResults[i])
			index = i
		}
		// fmt.Printf("turns : %d   steps : %d\n", len(allResults[i]), GetNumberOfSteps(allResults[i]))
	}
	// fmt.Println()
	return index
}

func Parser() (*Network, []string, string, string, int, error) {
	args := os.Args[1:]
	if len(args) != 1 {
		return nil, nil, "", "", 0, fmt.Errorf("ERROR: invalid data format, expected one argument (file name)")
	}
	file, err := os.Open(args[0])
	if err != nil {
		return nil, nil, "", "", 0, fmt.Errorf("ERROR: invalid data format, failed to open file: %v ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	if len(text) == 0 {
		return nil, nil, "", "", 0, fmt.Errorf("ERROR: invalid data format, file is empty")
	}
	NumberOfAnts, err := strconv.Atoi(text[0])
	if err != nil {
		return nil, nil, "", "", 0, fmt.Errorf("ERROR: invalid data format, invalid ant count: %s", text[0])
	}
	if NumberOfAnts < 1 {
		return nil, nil, "", "", 0, fmt.Errorf("ERROR: invalid data format, invalid number of Ants")
	}
	graph := &Network{}
	var Start, End string
	for i, line := range text {
		if line[0] != '#' && line[0] != 'L' {
			if strings.Contains(line, " ") {
				room := strings.Split(line, " ")
				if len(room) != 3 {
					return nil, nil, "", "", 0, fmt.Errorf("ERROR: invalid data format, invalid room format: %s", line)
				}
				_, err := strconv.Atoi(room[1])
				if err != nil {
					return nil, nil, "", "", 0, fmt.Errorf("ERROR: invalid data format, invalid room coordinates: %s", line)
				}
				_, err = strconv.Atoi(room[2])
				if err != nil {
					fmt.Println("ERROR: invalid data format, invalid room coordinates: " + line)
					return nil, nil, "", "", 0, fmt.Errorf("ERROR: invalid data format, invalid room coordinates: %s", line)
				}
				err = graph.AddVertex(room[0])
				if err != nil {
					return nil, nil, "", "", 0, err
				}
				if i > 0 && text[i-1] == "##start" {
					Start = room[0]
				} else if i > 0 && text[i-1] == "##end" {
					End = room[0]
				}
			} else if strings.Contains(line, "-") {
				edge := strings.Split(line, "-")
				if len(edge) != 2 {
					return nil, nil, "", "", 0, fmt.Errorf("ERROR: invalid data format, invalid tunnel format: %s", line)
				}
				err := graph.AddEdge(edge[0], edge[1])
				if err != nil {
					return nil, nil, "", "", 0, err
				}
			}
		} else if line[0] != '#' {
			return nil, nil, "", "", 0, fmt.Errorf("ERROR: invalid data format, room shouldn't start with L or #: %s", line)
		}
	}
	if Start == "" {
		return nil, nil, "", "", 0, fmt.Errorf("ERROR: invalid data format, missing start room")
	}
	if End == "" {
		return nil, nil, "", "", 0, fmt.Errorf("ERROR: invalid data format, missing end room")
	}
	return graph, text, Start, End, NumberOfAnts, nil
}
