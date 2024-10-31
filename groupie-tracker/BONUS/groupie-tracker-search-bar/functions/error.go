package groupie

import (
	"fmt"
	"net/http"
	"text/template"
)

type errorType struct {
	ErrorCode string
	Message   string
}

func ErrorPages(w http.ResponseWriter, code int) {
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		fmt.Println("Error parsing html file: ", err)
		return
	} else if code == 404 {
		err = t.Execute(w, errorType{ErrorCode: "404", Message: "Sorry, the page you are looking for does not exist."})
		if err != nil {
			t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		}
	} else if code == 405 {
		err = t.Execute(w, errorType{ErrorCode: "405", Message: "Method not allowed."})
		if err != nil {
			t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		}
	} else if code == 400 {
		err = t.Execute(w, errorType{ErrorCode: "400", Message: "Bad Request."})
		if err != nil {
			t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		}
	} else if code == 502 {
		err = t.Execute(w, errorType{ErrorCode: "502", Message: "Bad Gate way."})
		if err != nil {
			t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		}
	} else {
		t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
	}
}
