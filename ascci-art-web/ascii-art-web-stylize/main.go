package main

import (
	"fmt"
	"net/http"

	handler "docker/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("./template"))
	http.Handle("/template/*", http.StripPrefix("/template/", fs))

	http.HandleFunc("/", handler.Mainpage)
	http.HandleFunc("/ascii-art", handler.AsciiArt)
	http.HandleFunc("/ascii-art/Download", handler.DownloadFileHandler)
	fmt.Println("Staring Server at port http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println(err)
		return
	}
}