package main

import (
	"fmt"
	"os"

	fun "lem/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}
	containt, err := fun.Readfile(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	// fmt.Println(containt + "\n")

	farm, err := fun.ParseInput(containt)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	paths := fun.RetrievePaths(farm)

	uniquePaths := fun.Logic(paths)
	for _, v := range uniquePaths {
		fmt.Println(v)
	}

	if len(uniquePaths) == 0 {
		fmt.Println("ERROR: no valid path found")
		return
	}

	fun.Result(uniquePaths, farm.NumAnts)
}
