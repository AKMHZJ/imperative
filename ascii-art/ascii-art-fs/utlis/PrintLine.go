package ascii

import "fmt"

func PrintLine(words string, data []string) {
	for i := 0; i < 8; i++ {
		for _, char := range words {
			pos := StartIndex(char)
			fmt.Print(data[pos+i])
		}
		fmt.Println()
	}
}
