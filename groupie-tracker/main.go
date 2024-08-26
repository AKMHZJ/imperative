package main

import (
	"fmt"
	"log"
	"net/http"

	han "fetch/handlers"
)

func main() {
	http.HandleFunc("/", han.Home)
	http.HandleFunc("/Artist/{id}", han.ArtistHandler)

	fmt.Println("Starting Server At Port : http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
		return
	}
}
