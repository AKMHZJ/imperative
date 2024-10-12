package lem

func FindAllPaths(farm *Antfarm) [][]string {
	var allPaths [][]string
	visited := make(map[string]bool)
	var dfs func(current string, path []string)

	dfs = func(current string, path []string) {
		if current == farm.EndRoom.Name {
			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			allPaths = append(allPaths, pathCopy)
			return
		}

		visited[current] = true
		for _, neighbor := range farm.link[current] {
			if !visited[neighbor] {
				dfs(neighbor, append(path, neighbor))
			}
		}
		visited[current] = false
	}

	dfs(farm.StartRoom.Name, []string{farm.StartRoom.Name})
	return allPaths
}
