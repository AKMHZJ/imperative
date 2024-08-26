package docker

func PrintLine(words string, data []string) string {
	res := ""
	for i := 0; i < 8; i++ {
		for _, char := range words {
			pos := StartIndex(char)
			res += (data[pos+i])
		}
		res += "\n"
	}
	return res
}
