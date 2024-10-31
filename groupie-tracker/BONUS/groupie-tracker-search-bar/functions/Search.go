package groupie

import (
	"log"
	"net/http"
	"strings"
)

func Search(w http.ResponseWriter, r *http.Request) {
	input := strings.ToLower(r.FormValue("search"))
	if input == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data, err := FetchFullData()
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		ErrorPages(w, 500)
		return
	}

	var foundArtists []FullArtist
	for _, artist := range data {
		if matchesSearch(artist, input) {
			foundArtists = append(foundArtists, artist)
		}
	}

	if len(foundArtists) == 0 {
		http.Redirect(w, r, "/?error=no-results", http.StatusSeeOther)
		return
	}

	uniqueartists := uniqueArtists(foundArtists)
	uniqueData := uniqueArtists(data)

	if err := templates1.ExecuteTemplate(w, "index.html", struct {
		Search []FullArtist
		All    []FullArtist
	}{
		Search: uniqueartists,
		All:    uniqueData,
	}); err != nil {
		// log.Printf("Error executing template: %v", err)
		ErrorPages(w, 500)
		return
	}
}
