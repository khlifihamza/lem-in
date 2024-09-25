package functions

import (
	"fmt"
	"lem-in/entities"
	"strings"
)

func SimulateAntMovement(g *Graphh, paths [][]string) {
	var ants []*entities.Ant
	for i := 1; i <= g.AntCount; i++ {
		ants = append(ants, &entities.Ant{Id: i, PathIndex: -1, Position: -1, Finished: false})
	}

	finished := 0

	for finished < g.AntCount {
		movements := []string{}
		occupiedRooms := make(map[string]bool)

		// Move existing ants
		for i := range ants {
			if !ants[i].Finished && ants[i].PathIndex != -1 && ants[i].Position < len(paths[ants[i].PathIndex])-1 {
				nextRoom := paths[ants[i].PathIndex][ants[i].Position+1]
				if !occupiedRooms[nextRoom] || nextRoom == g.End {
					ants[i].Position++
					movements = append(movements, fmt.Sprintf("L%d-%s", ants[i].Id, nextRoom))
					if nextRoom != g.End {
						occupiedRooms[nextRoom] = true
					}
					if nextRoom == g.End {
						ants[i].Finished = true
						finished++
					}
				}
			}
		}

		// Deploy new ants
		for i := range ants {
			if ants[i].PathIndex == -1 && ants[i].Position == -1 {
				for j, path := range paths {
					if len(ants)-len(path) == getNumOfFinishedAnts(ants)-2 && i == len(ants)-1 && path[1] != g.End && path[0] == g.Start && len(paths) == 2 {
						continue // Wait for better path
					}
					if !occupiedRooms[path[1]] {
						ants[i].PathIndex = j
						ants[i].Position = 1
						movements = append(movements, fmt.Sprintf("L%d-%s", ants[i].Id, path[1]))
						occupiedRooms[path[1]] = true
						if path[1] == g.End {
							ants[i].Finished = true
							finished++
						}
						break
					}
				}
			}
		}

		if len(movements) > 0 {
			fmt.Println(strings.Join(movements, " "))
		}
	}
}

func getNumOfFinishedAnts(ants []*entities.Ant) int {
	finished := 0
	for _, ant := range ants {
		if ant.Finished {
			finished++
		}
	}
	return finished
}
