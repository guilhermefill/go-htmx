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
	data := map[string]interface{}{
		"Title":   "Spooky",
		"Content": "Spooky Times",
	}
	tmpl.Execute(w, data)
}
