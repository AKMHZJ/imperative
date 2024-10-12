package lem

import "sort"

func OptimizeAntFlow(farm *Antfarm, paths [][]string) [][]string {
	if len(paths) == 0 {
		return nil
	}

	// Sort paths by length
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	var optimizedPaths [][]string
	antsLeft := farm.numberofant

	for antsLeft > 0 {
		for i := range paths {
			if i == len(optimizedPaths) {
				optimizedPaths = append(optimizedPaths, make([]string, 0))
			}
			if antsLeft > 0 {
				optimizedPaths[i] = append(optimizedPaths[i], paths[i]...)
				antsLeft--
			}
		}
	}

	return optimizedPaths
}
