package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/maksonviini/Go-Book-Management-System/pkg/config"
	"github.com/maksonviini/Go-Book-Management-System/pkg/controllers"
)

func main() {
	err := config.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/", controllers.Create)
	r.Get("/", controllers.GetAll)
	r.Put("/{id}", controllers.Update)
	r.Delete("/{id}", controllers.Delete)
	r.Get("/{id}", controllers.Get)

	http.ListenAndServe(fmt.Sprint(":%s", config.GetServerPort()), r)
}
