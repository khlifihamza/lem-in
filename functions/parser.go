package functions

import (
	"bufio"
	"fmt"
	"lem-in/entities"
	"os"
	"strconv"
	"strings"
)

func ParseInput() (*Graphh, []string, error) {
	args := os.Args[1:]
	if len(args) != 1 {
		return nil, nil, fmt.Errorf("ERROR: invalid data format, expected one argument (file name)")
	}

	file, err := os.Open(args[0])
	if err != nil {
		return nil, nil, fmt.Errorf("ERROR: invalid data format, failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 0 {
		return nil, nil, fmt.Errorf("ERROR: invalid data format, file is empty")
	}

	g := &Graphh{
		Rooms:   make(map[string]*entities.Room),
		Tunnels: make(map[string][]*entities.Edge),
	}

	antCount, err := strconv.Atoi(lines[0])
	if err != nil {
		return nil, nil, fmt.Errorf("ERROR: invalid data format, invalid ant count: %s", lines[0])
	}
	if antCount <= 0 {
		return nil, nil, fmt.Errorf("ERROR: invalid data format, invalid number of Ants")
	}
	g.AntCount = antCount

	for i, line := range lines[1:] {
		line = strings.TrimSpace(line)
		if line == "" || line[0] == 'L' {
			continue
		}

		if line[0] == '#' {
			if line == "##start" || line == "##end" {
				if i+1 >= len(lines[1:]) {
					return nil, nil, fmt.Errorf("ERROR: invalid data format, missing room definition after %s", line)
				}
				nextLine := strings.TrimSpace(lines[i+2])
				if err := parseRoom(nextLine, g); err != nil {
					return nil, nil, err
				}
				if line == "##start" {
					g.Start = strings.Fields(nextLine)[0]
				} else {
					g.End = strings.Fields(nextLine)[0]
				}
			}
			continue
		}

		if strings.Contains(line, " ") {
			if err := parseRoom(line, g); err != nil {
				return nil, nil, err
			}
		} else if strings.Contains(line, "-") {
			if err := parseTunnel(line, g); err != nil {
				return nil, nil, err
			}
		} else {
			return nil, nil, fmt.Errorf("ERROR: invalid data format, unrecognized line format: %s", line)
		}
	}

	if g.Start == "" {
		return nil, nil, fmt.Errorf("ERROR: invalid data format, missing start room")
	}
	if g.End == "" {
		return nil, nil, fmt.Errorf("ERROR: invalid data format, missing end room")
	}

	return g, lines, nil
}

func parseRoom(line string, g *Graphh) error {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return fmt.Errorf("ERROR: invalid data format, invalid room format: %s", line)
	}

	name := parts[0]
	_, err1 := strconv.Atoi(parts[1])
	_, err2 := strconv.Atoi(parts[2])

	if err1 != nil || err2 != nil {
		return fmt.Errorf("ERROR: invalid data format, invalid room coordinates: %s", line)
	}

	g.AddRoom(name)
	return nil
}

func parseTunnel(line string, g *Graphh) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("ERROR: invalid data format, invalid tunnel format: %s", line)
	}

	room1, room2 := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	g.AddRoom(room1)
	g.AddRoom(room2)
	g.AddTunnel(room1, room2)

	return nil
}
