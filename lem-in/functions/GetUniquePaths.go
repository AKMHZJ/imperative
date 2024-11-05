package lem

import (
	"sort"
)


func Sort(allpaths [][]string) [][]string {
	sort.Slice(allpaths, func(i, j int) bool {
		return len(allpaths[i]) < len(allpaths[j])
	})
	return allpaths
}


func Logic(paths [][]string) [][]string {
	paths = Sort(paths)
	var allGroups [][][]string

	for _, path := range paths {
		tmpArr := [][]string{path}
		for _, otherPath := range paths {
			// Use helper function to check if path and otherPath are identical
			if areSlicesEqual(path, otherPath) {
				continue
			}

			// Check uniqueness with each path in the current group
			unique := true
			for _, groupPath := range tmpArr {
				if !checkIsUnique(otherPath, groupPath) {
					unique = false
					break
				}
			}
			if unique {
				tmpArr = append(tmpArr, otherPath)
			}
		}
		allGroups = append(allGroups, tmpArr)
	}

	// Find the largest group with the most paths and return it
	largestGroup := findLargestGroup(allGroups)
	return largestGroup
}


// areSlicesEqual compares two slices for equality
func areSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// findLargestGroup returns the group with the most paths
func findLargestGroup(groups [][][]string) [][]string {
	largestGroup := [][]string{}
	for _, group := range groups {
		if len(group) > len(largestGroup) {
			largestGroup = group
		}
	}
	return largestGroup
}
