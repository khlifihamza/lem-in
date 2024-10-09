package functions

import (
	"fmt"

	"lem-in/entities"
)

func DeployAntInCombination(Colony *Colony, paths [][]string) [][]string {
	var ants []*entities.Ant
	results := [][]string{}
	for i := 1; i <= Colony.NumberOfAnts; i++ {
		ants = append(ants, &entities.Ant{Id: i, PathIndex: -1, Position: -1, Finished: false})
	}
	finished := 0
	pathLimits := calculatePathLimits(paths, Colony.NumberOfAnts)
	for finished < Colony.NumberOfAnts {
		// fmt.Println(pathLimits)
		movements := []string{}
		occupiedRooms := make(map[string]bool)
		for i := range ants {
			if !ants[i].Finished && ants[i].PathIndex != -1 && ants[i].Position < len(paths[ants[i].PathIndex])-1 {
				nextRoom := paths[ants[i].PathIndex][ants[i].Position+1]
				if !occupiedRooms[nextRoom] || nextRoom == Colony.End {
					ants[i].Position++
					movements = append(movements, fmt.Sprintf("L%d-%s", ants[i].Id, nextRoom))
					if nextRoom != Colony.End {
						occupiedRooms[nextRoom] = true
					}
					if nextRoom == Colony.End {
						ants[i].Finished = true
						finished++
					}
				}
			}
		}
		for i := range ants {
			if ants[i].PathIndex == -1 && ants[i].Position == -1 {
				for j, path := range paths {
					if pathLimits[j] > 0 && !occupiedRooms[path[1]] {
						ants[i].PathIndex = j
						ants[i].Position = 1
						movements = append(movements, fmt.Sprintf("L%d-%s", ants[i].Id, path[1]))
						occupiedRooms[path[1]] = true
						pathLimits[j]--
						if path[1] == Colony.End {
							ants[i].Finished = true
							finished++
						}
						break
					}
				}
			}
		}
		if len(movements) > 0 {
			results = append(results, movements)
		}
	}
	return results
}

func DeployAntArmy(pathCombinations map[int][][]string, Colony *Colony) [][]string {
	allResults := map[int][][]string{}
	for key, pathpathCombination := range pathCombinations {
		for i, path := range pathpathCombination {
			newPath := []string{Colony.Start}
			newPath = append(newPath, path...)
			pathpathCombination[i] = newPath
		}
		pathpathCombination = Sort(pathpathCombination)
		// fmt.Println(pathpathCombination)
		allResults[key] = DeployAntInCombination(Colony, pathpathCombination)
	}
	index := CompareAllResults(allResults)
	// fmt.Println(pathCombinations[index])
	return allResults[index]
}
