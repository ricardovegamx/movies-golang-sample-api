package app

import (
	"github.com/ricardovegamx/movies-golang-sample-api/src/api/http/controllers"
)

func UrlMappings() {
	router.GET("/movies", controllers.Index)
}
