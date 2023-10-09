package api

import (
	"net/http"
	"text/template"
)

func Spooky(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/html/partials/spooky.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
