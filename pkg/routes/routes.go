package routes

import (
	"github.com/gorilla/mux"
	"github.com/guilhermefill/go-htmx/pkg/pages"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/", pages.Home).Methods("GET")
	router.HandleFunc("/about", pages.About).Methods("GET")
}
