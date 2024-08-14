package ascii

import "strings"

func PrintAscii(input string, data []string) string {
	output := ""
	lines  := strings.Count(input, "\\n")
	words  := strings.Split(input, "\\n")
	for j  := 0; j < len(words); j++ {
		
		// Handle new line
		if words[j] == "" {
			if lines > 0 {
				output += "\n"
				lines--
			}
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range words[j] {
				pos := StartIndex(char)
				output += data[pos+i]
			}
			output += "\n"
		}
	}
	return output
}
