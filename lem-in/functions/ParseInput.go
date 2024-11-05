package lem

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseInput(content string) (*AntFarm, error) {
	farm := &AntFarm{
		Rooms:       make(map[string]*Room),
		Connections: make(map[string][]string),
	}

	lines := strings.Split(content, "\n")
	lineNum := 0
	var currentCommand string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		lineNum++

		if line == "" || strings.HasPrefix(line, "#") {
			if line == "##start" {
				currentCommand += "start"
			} else if line == "##end" {
				currentCommand += "end"
			}
			continue
		}

		if lineNum == 1 {
			ants, err := strconv.Atoi(line)
			if err != nil || ants < 1 {
				return nil, fmt.Errorf("invalid ant number on line %d: %s", lineNum, line)
			}
			farm.NumAnts = ants

		} else if !strings.Contains(line, "-") {
			parts := strings.Fields(line)
			if len(parts) < 1 {
				return nil, fmt.Errorf("invalid room format on line %d: %s", lineNum, line)
			}

			name := parts[0]
			room := &Room{Name: name}

			if currentCommand == "start" {
				room.Isstart = true
				farm.Start = room
				currentCommand = ""
			} else if currentCommand == "end" {
				room.Isend = true
				farm.End = room
				currentCommand = ""
			}

			farm.Rooms[name] = room

		} else if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid link format on line %d: %s", lineNum, line)
			}

			room1, room2 := parts[0], parts[1]

			if _, exists := farm.Rooms[room1]; !exists {
				return nil, fmt.Errorf("unknown room in link on line %d: %s", lineNum, room1)
			}
			if _, exists := farm.Rooms[room2]; !exists {
				return nil, fmt.Errorf("unknown room in link on line %d: %s", lineNum, room2)
			}

			farm.Connections[room1] = append(farm.Connections[room1], room2)
			farm.Connections[room2] = append(farm.Connections[room2], room1)

		} else {
			return nil, fmt.Errorf("invalid line format on line %d: %s", lineNum, line)
		}
	}

	if farm.Start == nil || farm.End == nil {
		return nil, fmt.Errorf("start or end room not defined")
	}

	return farm, nil
}
