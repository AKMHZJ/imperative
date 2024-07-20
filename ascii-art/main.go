package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		fmt.Println("Usage : go run . [STRING]")
		return
	}

	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error")
		return
	}

	splitted := strings.Split(string(data), "\n")

	input := arg[0]
	lines := strings.Count(input, "\\n")
	words := strings.Split(input, "\\n")

	for j := 0; j < len(words); j++ {
		if words[j] == "" {
			if lines > 0 {
				fmt.Println()
				lines--
			}
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range words[j] {
				pos := (char - 32) + 1 + 8*(char-32)
				fmt.Print(splitted[int(pos)+i])
			}
			fmt.Println()

		}

	}
}
