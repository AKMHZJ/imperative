package lem

import (
	"fmt"
	"strings"
)

func MoveAntsOptimized(farm *Antfarm, paths [][]string) {
	if len(paths) == 0 {
		return
	}

	antMoves := make([][]string, len(paths[0]))
	antNum := 1

	for i := range paths[0] {
		for j := range paths {
			if i < len(paths[j]) {
				if i > 0 || j == 0 {
					move := fmt.Sprintf("L%d-%s", antNum, paths[j][i])
					antMoves[i] = append(antMoves[i], move)
					antNum++
				}
			}
		}
	}

	for _, moves := range antMoves {
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}
