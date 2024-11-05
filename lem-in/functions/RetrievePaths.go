package lem

func RetrievePaths(farm *AntFarm) [][]string {
	return TraversePath(farm.Start.Name, farm.End.Name, []string{}, farm.Connections, [][]string{})
}
