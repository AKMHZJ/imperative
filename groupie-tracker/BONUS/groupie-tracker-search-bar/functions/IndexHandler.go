package groupie

import (
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPages(w, 404)
		return
	}

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

	uniqueData := uniqueArtists(data)

	if err := templates1.ExecuteTemplate(w, "index.html", struct {
		Search []FullArtist
		All    []FullArtist
	}{
		Search: uniqueData,
		All:    uniqueData,
	}); err != nil {
		// log.Printf("Error executing template: %v", err)
		ErrorPages(w, 500)
		return
	}
}
