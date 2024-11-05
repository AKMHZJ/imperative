package lem

import (
	"fmt"
	"sort"
)

func Result(ShortestPaths [][]string, AntsNumber int) {
	sort.Slice(ShortestPaths, func(i, j int) bool {
		return len(ShortestPaths[i]) < len(ShortestPaths[j])
	})
	AntAndPath := make([][]int, len(ShortestPaths))
	LenPaths := len(ShortestPaths)
	i := 1
	for i <= AntsNumber {
		for j := 0; j < LenPaths; j++ {
			if LenPaths == 1 || j == LenPaths-1 && len(ShortestPaths[j-1])+len(AntAndPath[j-1]) > len(ShortestPaths[j])+len(AntAndPath[j]) || j != LenPaths-1 && len(ShortestPaths[j])+len(AntAndPath[j]) <= len(ShortestPaths[j+1])+len(AntAndPath[j+1]) {
				AntAndPath[j] = append(AntAndPath[j], i)
				i++
				if i > AntsNumber {
					break
				}
			} else {
				break
			}
		}
	}
	i = 1
	curr := 1
	Position := map[int]int{}
	for i <= AntsNumber {
		for j := 0; j < LenPaths; j++ {
			for t := 0; t < curr && t < len(AntAndPath[j]); t++ {
				Position[AntAndPath[j][t]]++
				if Position[AntAndPath[j][t]] < len(ShortestPaths[j]) {
					fmt.Print("L", AntAndPath[j][t], "-", ShortestPaths[j][Position[AntAndPath[j][t]]])
					fmt.Print(" ")
					if Position[AntAndPath[j][t]] == len(ShortestPaths[j])-1 {
						i++
					}
				}
			}
		}
		fmt.Println()
		curr++
	}
}
