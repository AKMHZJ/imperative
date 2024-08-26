package docker

func StartIndex(c rune) int {
	return int((c - 32) + 1 + 8*(c-32))
}
