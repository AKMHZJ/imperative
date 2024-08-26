package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Send(id string) Artist {
	url := "https://groupietrackers.herokuapp.com/api/"
	urlartist := url + "artists/" + id
	var artist Artist
	err := getwithId(urlartist, &artist)
	if err != nil {
		log.Fatal(fmt.Errorf("error fetching %v", err))
	}

	urllocations := url + "locations/" + id
	var location Location
	err = getwithId(urllocations, &location)
	if err != nil {
		log.Fatal(fmt.Errorf("error fetching %v", err))
	}

	urldates := url + "dates/" + id
	var dates Date
	err = getwithId(urldates, &dates)
	if err != nil {
		log.Fatal(fmt.Errorf("error fetching %v", err))
	}

	urlrelation := url + "relation/" + id
	var relation Relation
	err = getwithId(urlrelation, &relation)
	if err != nil {
		log.Fatal(fmt.Errorf("error fetching %v", err))
	}
	artist.Locations = location.Locations
	artist.Dates = dates.Dates
	artist.DatesLocations = relation.DatesLocations
	return artist
}

func getwithId(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error %v", err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(fmt.Errorf("error fetching %v", err))
	}
	return nil
}

func FetchData() (ApiData, error) {
	baseURL := "https://groupietrackers.herokuapp.com/api"
	endpoints := []string{"artists", "locations", "dates", "relation"}

	apiData := ApiData{}

	for _, endpoint := range endpoints {
		url := fmt.Sprintf("%s/%s", baseURL, endpoint)

		resp, err := http.Get(url)
		if err != nil {
			return apiData, fmt.Errorf("error fetching %s: %v", endpoint, err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return apiData, fmt.Errorf("error reading %s response: %v", endpoint, err)
		}

		switch endpoint {
		case "artists":
			err = json.Unmarshal(body, &apiData.Artists)
		case "locations":
			err = json.Unmarshal(body, &apiData.Locations)
		case "dates":
			err = json.Unmarshal(body, &apiData.Dates)
		case "relation":
			err = json.Unmarshal(body, &apiData.Relations)
		}

		if err != nil {
			return apiData, fmt.Errorf("error unmarshaling %s data: %v", endpoint, err)
		}
	}

	return apiData, nil
}

