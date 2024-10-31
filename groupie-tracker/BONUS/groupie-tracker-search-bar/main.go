package main

import (
	"fmt"
	"log"
	"net/http"

	link "groupie/functions"
)

func main() {
	http.HandleFunc("/", link.IndexHandler)
	http.HandleFunc("/result/", link.ResultHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/search/", link.Search)

	fmt.Println("\033[32mServer started at http://localhost:8080\033[0m")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
