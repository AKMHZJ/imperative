package lem

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseContent(content string) (*Antfarm, error) {
	farm := &Antfarm{
		Rooms: make(map[string]*Room),
		link:  make(map[string][]string),
	}

	lines := strings.Split(content, "\n")
	lineNum := 0
	var currentcommand string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		lineNum++

		if line == "" || strings.HasPrefix(line, "#") {
			if line == "##start" {
				currentcommand += "start"
			} else if line == "##end" {
				currentcommand += "end"
			}
			continue
		}

		if lineNum == 1 {
			ants, err := strconv.Atoi(line)
			if err != nil || ants < 1 {
				return nil, fmt.Errorf("invalid ant number in line %d: %s", lineNum, line)
			}
			farm.numberofant = ants
		} else if !strings.Contains(line, "-") {
			parts := strings.Fields(line)
			if len(parts) < 1 {
				return nil, fmt.Errorf("invalid room format on line %d: %s", lineNum, line)
			}

			name := parts[0]
			room := &Room{Name: name}

			if currentcommand == "start" {
				room.Isstart = true
				farm.StartRoom = room
				currentcommand = ""
			} else if currentcommand == "end" {
				room.Isend = true
				farm.EndRoom = room
				currentcommand = ""
			}
			farm.Rooms[name] = room
		} else if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid link format on line %d: %s", lineNum, line)
			}
			room1, room2 := parts[0], parts[1]
			if _, exist := farm.Rooms[room1]; !exist {
				return nil, fmt.Errorf("unknown room in link on line %d: %s", lineNum, room1)
			}
			if _, exist := farm.Rooms[room2]; !exist {
				return nil, fmt.Errorf("unknown room in link on line %d: %s", lineNum, room2)
			}
			farm.link[room1] = append(farm.link[room1], room2)
			farm.link[room2] = append(farm.link[room2], room1)
		} else {
			return nil, fmt.Errorf("invalid line format on line %d: %s", lineNum, line)
		}
	}
	if farm.StartRoom == nil || farm.EndRoom == nil {
		return nil, fmt.Errorf("start or end room not found")
	}
	return farm, nil
}
