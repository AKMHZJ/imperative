package ascii

func IsValidArgs(args []string) bool {
	return len(args) != 0 && len(args) <= 3
}
