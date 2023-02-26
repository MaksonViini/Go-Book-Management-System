package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/maksonviini/Go-Book-Management-System/pkg/config"
	"github.com/maksonviini/Go-Book-Management-System/pkg/controllers"
)

var RegisterBookStoreRoutesfunc = func(router *chi.Mux) {
	err := config.Load()

	if err != nil {
		panic(err)
	}

	router.Post("/", controllers.Create)
	router.Get("/", controllers.GetAll)
	router.Put("/{id}", controllers.Update)
	router.Delete("/{id}", controllers.Delete)
	router.Get("/{id}", controllers.Get)

}
