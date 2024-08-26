package fetch

import (
	api "fetch/functions"
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "405 Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	var data api.ApiData
	data, _ = api.FetchData()

	tmp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 internal server error.", http.StatusInternalServerError)
		return
	}
	err = tmp.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 internal server error.", http.StatusInternalServerError)
		return
	}
}
