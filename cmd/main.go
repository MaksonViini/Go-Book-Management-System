package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/maksonviini/Go-Book-Management-System/pkg/config"
	"github.com/maksonviini/Go-Book-Management-System/pkg/routes"
)

func main() {
	router := gin.Default()

	routes.RegisterBookStoreRoutesfunc(&router.RouterGroup)

	http.Handle("/", router)
	log.Println("Listening...")

	if err := router.Run(fmt.Sprintf(":%s", config.GetServerPort())); err != nil {
		log.Fatal(err)
	}
}
