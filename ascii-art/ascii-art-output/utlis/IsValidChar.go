package ascii

func IsValidChar(r rune) bool {
	return r >= ' ' && r <= '~'
}
