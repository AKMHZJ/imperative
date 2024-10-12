package lem

import (
	"fmt"
	"strings"
)

func MoveAnts(farm *Antfarm, path []string) {
	if len(path) < 2 {
		return
	}

	type Ant struct {
		id       int
		position int
	}

	ants := make([]Ant, 0, farm.numberofant)
	nextAntID := 1
	finished := 0

	for finished < farm.numberofant {
		moves := []string{}

		// Move existing ants
		for i := range ants {
			if ants[i].position < len(path)-1 {
				ants[i].position++
				moves = append(moves, fmt.Sprintf("L%d-%s", ants[i].id, path[ants[i].position]))

				if ants[i].position == len(path)-1 {
					finished++
				}
			}
		}

		// Add new ant if there's space
		if len(ants) < len(path)-1 && nextAntID <= farm.numberofant {
			ants = append(ants, Ant{id: nextAntID, position: 1})
			moves = append(moves, fmt.Sprintf("L%d-%s", nextAntID, path[1]))
			nextAntID++
		}

		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}

		// Remove ants that have finished
		newAnts := ants[:0]
		for _, ant := range ants {
			if ant.position < len(path)-1 {
				newAnts = append(newAnts, ant)
			}
		}
		ants = newAnts
	}
}
