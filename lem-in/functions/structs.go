package lem

type Room struct {
	Name    string
	Isstart bool
	Isend   bool
}

type AntFarm struct {
	NumAnts     int
	Rooms       map[string]*Room
	Connections map[string][]string
	Start       *Room
	End         *Room
}
