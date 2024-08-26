package docker

import (
	utils "docker/utlis"
	"net/http"
	"text/template"
)

type MyStruct struct {
    Data    string
    Banners string
	Result string
}

func AsciiArt(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "405 Method not allowed.", http.StatusMethodNotAllowed)
		return
	}	

	data :=  r.FormValue("text")
	banners := r.FormValue("banner")	
    
	db := MyStruct{
    Data    : data ,
    Banners : banners,
	}

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

	tmp, err := template.ParseFiles("template/download.html")
	if err != nil {
		http.Error(w, "500 internal server error.", http.StatusInternalServerError)
		return
	}
	db.Result = result
	err = tmp.Execute(w,db)
	if err != nil {
		http.Error(w, "500 internal server error.", http.StatusInternalServerError)
		return
	}

}
