package app

import (
	"github.com/ricardovegamx/movies-golang-sample-api/src/api/http/controllers"
)

func UrlMappings() {
	router.GET("/movies", controllers.Index)
	router.GET("/movies/:movie_id", controllers.Get)
	router.PUT("/movies/:movie_id", controllers.Update)
	router.POST("/movies", controllers.Create)
	router.DELETE("/movies/:movie_id", controllers.Delete)
}
