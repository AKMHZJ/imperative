package docker

func IsValidChar(r rune) bool {
	return (r >= ' ' && r <= '~') || (r == '\r' || r == '\n')
}
