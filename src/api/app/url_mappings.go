package app

import (
	"github.com/ricardovegamx/movies-api/src/api/http/controllers"
)

func UrlMappings() {
	router.GET("/movies", controllers.Index)
}
