package lem

func FindShortestPath(farm *Antfarm) []string {
	visited := make(map[string]bool)
	queue := [][2]interface{}{{farm.StartRoom.Name, []string{farm.StartRoom.Name}}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		roomName := current[0].(string)
		path := current[1].([]string)

		if roomName == farm.EndRoom.Name {
			return path
		}

		if !visited[roomName] {
			visited[roomName] = true
			for _, neighbor := range farm.link[roomName] {
				if !visited[neighbor] {
					newPath := make([]string, len(path))
					copy(newPath, path)
					newPath = append(newPath, neighbor)
					queue = append(queue, [2]interface{}{neighbor, newPath})
				}
			}
		}
	}
	return nil
}