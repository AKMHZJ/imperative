package lem

func IsRoomVisited(path []string, room string) bool {
	for _, p := range path {
		if p == room {
			return true
		}
	}
	return false
}
