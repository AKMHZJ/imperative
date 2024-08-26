package fetch

import (
	"bytes"
	"html/template"
	"net/http"

	api "fetch/functions"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	IdNmber := r.PathValue("id")
	data := api.Send(IdNmber)
	tmp, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		http.Error(w, "500 internal server error.", http.StatusInternalServerError)
		return
	}
	buf := bytes.Buffer{}
	err = tmp.Execute(&buf, data)
	if err != nil {
		http.Error(w, "500 internal server error.", http.StatusInternalServerError)
		return
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
}
