package main

import (
	"fmt"
	"os"

	utils "ascii/utlis"
)

func main() {
	args := os.Args[1:]

	if !utils.IsValidArgs(args) {
		fmt.Println("usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	Args, err := utils.HandleArgs(args)
	utils.CheckError(err)

	input := Args.Input
	data, err := utils.HandleBanner(Args.Banner)
	utils.CheckError(err)

	err = utils.IsValidInput(input, data)
	utils.CheckError(err)

	output := utils.PrintAscii(input, data)

	if Args.Output == "" {
		fmt.Print(output)
	} else {
		err := utils.WriteToFile(Args.Output, output)
		utils.CheckError(err)
	}
}
