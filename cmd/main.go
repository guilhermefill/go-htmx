package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/guilhermefill/go-htmx/pkg/routes"
)

func main() {
	fmt.Println("Starting server on port 4200...")

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)

	srv := &http.Server{
		Addr:         ":4200",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
