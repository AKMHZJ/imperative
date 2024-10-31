package groupie

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func ResultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorPages(w, 405)
		return
	}

	data, err := FetchFullData()
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		ErrorPages(w, 500)
		return
	}

	IdNmber, err := strconv.Atoi(r.URL.Query().Get("id"))
	if IdNmber <= 0 || IdNmber > len(data) || err != nil {
		ErrorPages(w, 400)
		return
	}

	ApiData.ArtistData = data[IdNmber-1].Artists

	ApiData.Locations.Locations = data[IdNmber-1].Locations

	response2, err := http.Get(Api + "/dates/" + r.URL.Query().Get("id"))
	if err != nil {
		ErrorPages(w, 502)
		return
	}
	defer response2.Body.Close()
	var stockDates Dates
	err = json.NewDecoder(response2.Body).Decode(&stockDates)
	if err != nil {
		ErrorPages(w, 500)
		return
	}
	ApiData.Dates = stockDates

	response4, err := http.Get(Api + "/relation/" + r.URL.Query().Get("id"))
	if err != nil {
		ErrorPages(w, 502)
		return
	}
	defer response4.Body.Close()
	var stockRelation Relations
	err = json.NewDecoder(response4.Body).Decode(&stockRelation)
	if err != nil {
		ErrorPages(w, 500)
		return
	}
	ApiData.Relation = stockRelation

	if err := templates2.ExecuteTemplate(w, "result.html", ApiData); err != nil {
		// log.Printf("Error executing template: %v", err)
		ErrorPages(w, 500)
		return
	}
}
