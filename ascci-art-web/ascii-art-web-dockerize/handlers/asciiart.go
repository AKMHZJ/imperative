package docker

import (
	utils "docker/utlis"
	"fmt"
	"net/http"
)

func AsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	data := r.FormValue("text")
	banners := r.FormValue("banner")

	if len(data) == 0 {
		http.Error(w, "400 bad request.", http.StatusBadRequest)
		return
	}

	if len(data) > 1000 {
		http.Error(w, "413 Request Entity Too Large.", http.StatusRequestEntityTooLarge)
		return
	}

	result, err := utils.Generator(data, banners)
	if err != nil {
		http.Error(w, "400 bad request.", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s", result)
}
