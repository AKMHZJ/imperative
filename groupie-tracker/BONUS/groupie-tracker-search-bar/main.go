package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	data PageData
	tmp  *template.Template
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type APIindex struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
}

type APIlocations struct {
	Index []struct {
		Location []string `json:"locations"`
	} `json:"index"`
}

type APIdates struct {
	Index []struct {
		Dates []string `json:"dates"`
	} `json:"index"`
}

type APIrelations struct {
	Index []struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type Card struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	ConcertDates []string
	Relation     map[string][]string
}

type PageData struct {
	Cards         []Card
	ArtistNames   map[string]bool
	Members       map[string]bool
	CreationDates map[string]bool
	FirstAlbums   map[string]bool
	Locations     map[string]bool
}

func init() {
	tmp = template.Must(template.ParseGlob("templates/*.html"))
}

func FetchData(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	return json.Unmarshal(body, target)
}

func fetchartistinfo(url string) ([]Card, error) {
	var api APIindex
	if err := FetchData(url, &api); err != nil {
		return nil, err
	}
	var artists []Artist
	if err := FetchData(api.Artists, &artists); err != nil {
		return nil, err
	}
	var location APIlocations
	if err := FetchData(api.Locations, &location); err != nil {
		return nil, err
	}
	var dates APIdates
	if err := FetchData(api.Dates, &dates); err != nil {
		return nil, err
	}
	var relation APIrelations
	if err := FetchData(api.Relations, &relation); err != nil {
		return nil, err
	}
	var cards []Card
	for i, artist := range artists {
		cards = append(cards, Card{
			Id:           artist.Id,
			Image:        artist.Image,
			Members:      artist.Members,
			CreationDate: artist.CreationDate,
			FirstAlbum:   artist.FirstAlbum,
			Locations:    location.Index[i].Location,
			ConcertDates: dates.Index[i].Dates,
			Relation:     relation.Index[i].DatesLocations,
		})
	}
	return cards, nil
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != "GET" {
		http.Error(w, "TH PAGE ARE YOU LOOKING FOR NOT FOUND LOOL !!", 404)
		return
	}
	API := "https://groupietrackers.herokuapp.com/api"

	LISTS, err := fetchartistinfo(API)
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR LOL !! ", http.StatusInternalServerError)
		log.Printf("Error fetching artist data ! %v", err)
		return
	}

	artistNames, members, creationDates, firstAlbums, locations := uniqueValues(LISTS)

	data = PageData{
		Cards:         LISTS,
		ArtistNames:   artistNames,
		Members:       members,
		CreationDates: creationDates,
		FirstAlbums:   firstAlbums,
		Locations:     locations,
	}
	if err := tmp.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func uniqueValues(cards []Card) (map[string]bool, map[string]bool, map[string]bool, map[string]bool, map[string]bool) {
	artistNames := make(map[string]bool)
	members := make(map[string]bool)
	creationDates := make(map[string]bool)
	firstAlbums := make(map[string]bool)
	locations := make(map[string]bool)

	for _, card := range cards {
		artistNames[card.Name] = true
		for _, member := range card.Members {
			members[member] = true
		}
		creationDates[strconv.Itoa(card.CreationDate)] = true
		firstAlbums[card.FirstAlbum] = true
		for _, location := range card.Locations {
			locations[location] = true
		}
	}

	return artistNames, members, creationDates, firstAlbums, locations
}

func SearchResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Bad Request: Only GET method is allowed", http.StatusBadRequest)
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Bad Request: No search query provided", http.StatusBadRequest)
		return
	}

	query = strings.ToLower(query)
	var results []Card

	for _, card := range data.Cards {
		// Check for name match
		if strings.Contains(strings.ToLower(card.Name), query) {
			results = append(results, card)
		} else if Containesitem(card.Members, query) {
			results = append(results, card)
		} else if strings.Contains(strings.ToLower(card.FirstAlbum), query) {
			results = append(results, card)
		} else if strconv.Itoa(card.CreationDate) == query {
			results = append(results, card)
		} else if Containesitem(card.Locations, query) {
			results = append(results, card)
		}
	}

	// Render the results to the template
	if err := tmp.ExecuteTemplate(w, "search.html", results); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func Containesitem(slice []string, query string) bool {
	for _, item := range slice {
		if strings.Contains(strings.ToLower(item), query) {
			return true
		}
	}
	return false
}

func Bandinfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Bad Request: Only GET method is allowed", http.StatusBadRequest)
	}
	id := strings.TrimPrefix(r.URL.RawQuery, "=id")
	if id == "" {
		http.Error(w, "/400", http.StatusBadRequest)
		return
	}
	uno, err := strconv.Atoi(id)
	if err != nil || uno > 52 || uno < 0 {
		http.Error(w, "this page is not found", 404)
		return
	}
	tmp.ExecuteTemplate(w, "bandsinfo.html", data.Cards[uno-1])
}

func main() {
	// first** lancer server for the first time
	http.HandleFunc("/", index)
	http.HandleFunc("/search", SearchResult)
	http.HandleFunc("/bandsinfo", Bandinfo)
	fmt.Println("Server lancer on port 8080 at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
