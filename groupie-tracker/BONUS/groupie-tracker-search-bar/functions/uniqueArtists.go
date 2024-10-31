package groupie

func uniqueArtists(data []FullArtist) []FullArtist {
	seen := make(map[string]bool)
	var uniqueLocations []string
	var unique []FullArtist

	for _, artist := range data {
		for _, location := range artist.Locations {
			if !seen[location] {
				seen[location] = true
				uniqueLocations = append(uniqueLocations, location)
			}
		}
		artist.Locations = uniqueLocations
		if !seen[artist.Artists.Name] {
			seen[artist.Artists.Name] = true
			unique = append(unique, artist)
			uniqueLocations = []string{}
		}
	}

	return unique
}
