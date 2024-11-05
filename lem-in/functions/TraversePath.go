package lem

func TraversePath(start, end string, path []string, graph map[string][]string, allPaths [][]string) [][]string {
	path = append(path, start)
	if start == end {
		finalPath := make([]string, len(path))
		copy(finalPath, path)
		return append(allPaths, finalPath)
	}
	for _, nextRoom := range graph[start] {
		if !IsRoomVisited(path, nextRoom) {
			allPaths = TraversePath(nextRoom, end, path, graph, allPaths)
		}
	}
	return allPaths
}
