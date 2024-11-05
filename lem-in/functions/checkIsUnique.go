package lem

// checkIsUnique checks if two paths have unique intermediate rooms (excluding start and end)
func checkIsUnique(arr, arr1 []string) bool {
	midArr := arr[1 : len(arr)-1]
	midArr1 := arr1[1 : len(arr1)-1]

	for _, element := range midArr {
		for _, element1 := range midArr1 {
			if element == element1 {
				return false
			}
		}
	}
	return true
}
