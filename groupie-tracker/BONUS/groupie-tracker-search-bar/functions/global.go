package groupie

import "text/template"

type ApiOfArtist struct {
	ArtistData ArtistData `json:"artists"`
	Locations  Locations  `json:"locations"`
	Dates      Dates      `json:"dates"`
	Relation   Relations  `json:"relation"`
}

type ArtistData struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type FullArtist struct {
	Artists   ArtistData
	Locations []string
}

const Api = "https://groupietrackers.herokuapp.com/api"

var DATA []ArtistData

var (
	templates1 = template.Must(template.ParseFiles("templates/index.html"))
	templates2 = template.Must(template.ParseFiles("templates/result.html"))
)

var ApiData ApiOfArtist
