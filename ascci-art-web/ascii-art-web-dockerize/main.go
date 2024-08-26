package main

import (
	"fmt"
	"net/http"

	handler "docker/handlers"
)

func main() {
	http.HandleFunc("/", handler.Mainpage)
	http.HandleFunc("/ascii-art", handler.AsciiArt)

	fmt.Println("Staring Server at port http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println(err)
		return
	}
}
