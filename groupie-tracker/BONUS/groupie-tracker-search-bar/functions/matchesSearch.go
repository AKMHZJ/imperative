package groupie

import (
	"strconv"
	"strings"
)

func matchesSearch(artist FullArtist, input string) bool {
	if strings.Contains(strings.ToLower(artist.Artists.Name), strings.ToLower(input)) ||
		strings.Contains(strings.ToLower(artist.Artists.FirstAlbum), strings.ToLower(input)) {
		return true
	}
	for _, member := range artist.Artists.Members {
		if strings.Contains(strings.ToLower(member), strings.ToLower(input)) {
			return true
		}
	}

	for _, location := range artist.Locations {
		if strings.Contains(strings.ToLower(location), strings.ToLower(input)) {
			return true
		}
	}

	if datem, err := strconv.Atoi(input); err == nil && datem == artist.Artists.CreationDate {
		return true
	}
	return false
}
