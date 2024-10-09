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

// Contains checks if a given key exists in a slice of Vertex pointers and returns true if found, otherwise false.
func Contains(slice []*entities.Vertex, key string) bool {
	for _, vertex := range slice {
		if key == vertex.Key {
			return true
		}
	}
	return false
}

// RemoveFromSlice removes a specified vertex from a slice of Vertex pointers and returns the updated slice.
func RemoveFromSlice(slice []*entities.Vertex, vertex *entities.Vertex) []*entities.Vertex {
	for i, v := range slice {
		if v == vertex {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// ContainsInslice checks if a given key exists in a slice of strings and returns true if found, otherwise false.
func ContainsInslice(slice []string, key string) bool {
	for _, vertex := range slice {
		if key == vertex {
			return true
		}
	}
	return false
}

// DeleteInSlice removes the slice at the specified index from the 2D slice shortestPaths and returns the new slice.
func DeleteInSlice(shortestPaths [][]string, index int) [][]string {
	newSlice := [][]string{}
	for i, slice := range shortestPaths {
		if i != index {
			newSlice = append(newSlice, slice)
		}
	}
	return newSlice
}

// Sort orders the 2D slice paths based on the length of each inner slice, arranging them in ascending order.
func Sort(paths [][]string) [][]string {
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
	return paths
}

// GetNumberOfSteps calculates the total number of steps by summing the lengths of each inner slice in the 2D slice result.
func GetNumberOfSteps(result [][]string) int {
	c := 0
	for _, slice := range result {
		c += len(slice)
	}
	return c
}

// PrintMovements prints the movements in a formatted way, separating each step with a space and adding a newline after each movement.
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

// CompareAllResults compares multiple results to find the index of the one with the least turns and steps.
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

// Parser reads input from a file and constructs a graph representation of rooms and tunnels.
// It returns a network of rooms, the lines read from the file, the start room,
// the end room, the number of ants, and any error encountered during parsing.
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
