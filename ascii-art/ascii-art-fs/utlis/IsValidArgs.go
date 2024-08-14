package ascii

func IsValidArgs(args []string) bool {
	return !(len(args) > 2 || len(args) == 0)
}
