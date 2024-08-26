package fetch

type Artist struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Locations      interface{}         `json:"locations"`
	Dates          []string            `json:"dates"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Location struct {
	Locations interface{} `json:"locations"`
}

type Date struct {
	Dates []string `json:"dates"`
}

type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type ApiData struct {
	Artists   []Artist
	Locations struct {
		Index []Location `json:"index"`
	}
	Dates struct {
		Index []Date `json:"index"`
	}
	Relations struct {
		Index []Relation `json:"index"`
	}
}

type Errors struct {
	Message string
	Code    int
	Desc string
}

type AllMes struct {
	NotFound string
	MethoNot string
}
