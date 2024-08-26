package docker

import (
	utils "docker/utlis"
	"fmt"
	"net/http"
)

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "405 Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	data :=  r.FormValue("t")
	banners := r.FormValue("b")

	result, err := utils.Generator(data, banners)
	if err != nil {
		http.Error(w, "400 bad request.", http.StatusBadRequest)
		return
	}

	fileName := "ascii-art-Web.txt"

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(result)))

	w.Write([]byte(result))

}
