package groupie

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func FetchFullData() ([]FullArtist, error) {
	result := make([]FullArtist, 0)
	artists := []ArtistData{}

	response, err := http.Get(Api + "/artists")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch artists: %w", err)
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&artists); err != nil {
		return nil, fmt.Errorf("failed to decode artists: %w", err)
	}

	for _, artist := range artists {
		locations, err := getLocationByID(strconv.Itoa(artist.ID))
		if err != nil {
			log.Printf("Warning: failed to get locations for artist %d: %v", artist.ID, err)
			continue
		}

		result = append(result, FullArtist{
			Artists:   artist,
			Locations: locations.Locations,
		})
	}

	return result, nil
}
