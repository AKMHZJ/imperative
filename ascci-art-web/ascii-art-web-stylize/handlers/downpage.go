package docker

import (
	"net/http"
	"strconv"

	utils "docker/utlis"
)

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	data := r.FormValue("t")
	banners := r.FormValue("b")

	result, err := utils.Generator(data, banners)
	if err != nil {
		http.Error(w, "400 bad request.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art-Web.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(result)))

	w.Write([]byte(result))
}
