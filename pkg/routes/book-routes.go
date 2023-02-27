package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maksonviini/Go-Book-Management-System/pkg/config"
	"github.com/maksonviini/Go-Book-Management-System/pkg/controllers"
)

var RegisterBookStoreRoutesfunc = func(router *gin.RouterGroup) {

	err := config.Load()

	if err != nil {
		panic(err)
	}

	router.POST("/", controllers.Create)
	router.GET("/", controllers.GetAll)
	router.PUT("/:id", controllers.Update)
	router.DELETE("/:id", controllers.Delete)
	router.GET("/:id", controllers.Get)

}
