package main

import (
	"fmt"
	"os"

	util "lem/functions"
)

// func moveAnts(farm *Antfarm, path []string) {
// 	if len(path) < 2 {
// 		return
// 	}

// 	antMoves := make([][]string, len(path)-1)
// 	for i := 0; i < farm.numberofant; i++ {
// 		Antnum := i + 1
// 		for j := 0; j < len(path)-1; j++ {
// 			if i-j >= 0 {
// 				move := fmt.Sprintf("L%d-%s", Antnum, path[j+1])
// 				antMoves[j] = append(antMoves[j], move)
// 			}
// 		}
// 	}

// 	for _, moves := range antMoves {
// 		fmt.Println((strings.Join(moves, " ")))
// 	}
// }

// ************************************
// ************************************

// ************************************
// ************************************

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run [Option] <filename>")
		return
	}

	filename := os.Args[1]
	content, err := util.Readfile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	farm, err := util.ParseContent(content)
	if err != nil {
		fmt.Printf("Error parsing content: %v\n", err)
		return
	}

	// fmt.Println(content)

	shortestpath := util.FindShortestPath((farm))
	if shortestpath == nil {
		fmt.Println("Error: No path found from start to end")
		return
	}

	util.MoveAnts(farm, shortestpath)

	allPaths := util.FindAllPaths(farm)
	if len(allPaths) == 0 {
		fmt.Println("Error: No path found from start to end")
		return
	}

	optimizedPaths := util.OptimizeAntFlow(farm, allPaths)

	util.MoveAntsOptimized(farm, optimizedPaths)
	// fmt.Printf("Number of ants: %d\n", farm.numberofant)
}
