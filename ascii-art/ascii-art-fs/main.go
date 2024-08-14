package main

import (
	"fmt"
	"os"
	"strings"

	utils "ascii/utlis"
)

func main() {
	args := os.Args[1:]

	if !utils.IsValidArgs(args) {
		fmt.Println("Usage : go run . [STRING] [BANNER]")
		return
	}
	input := args[0]
	if input == "" {
		return
	}
	data, err := utils.HandleBanner(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := utils.IsValidInput(input, data); err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Count(input, "\\n")
	words := strings.Split(input, "\\n")

	for j := 0; j < len(words); j++ {
		// Handle new line
		if words[j] == "" {
			if lines > 0 {
				fmt.Println()
				lines--
			}
			continue
		}
		utils.PrintLine(words[j], data)
	}
}
