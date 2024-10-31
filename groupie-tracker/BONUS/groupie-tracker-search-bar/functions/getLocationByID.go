package groupie

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getLocationByID(id string) (Locations, error) {
	var res Locations

	response, err := http.Get(Api + "/locations/" + id)
	if err != nil {
		return res, fmt.Errorf("failed to fetch locations: %w", err)
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
		return res, fmt.Errorf("failed to decode locations: %w", err)
	}
	return res, nil
}
