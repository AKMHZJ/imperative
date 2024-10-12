package lem

type Room struct {
	Name    string
	Isstart bool
	Isend   bool
}

type Antfarm struct {
	numberofant int
	Rooms       map[string]*Room
	link        map[string][]string
	StartRoom   *Room
	EndRoom     *Room
}
