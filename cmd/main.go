package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/maksonviini/Go-Book-Management-System/pkg/config"
	"github.com/maksonviini/Go-Book-Management-System/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	routes.RegisterBookStoreRoutesfunc(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r))
}
