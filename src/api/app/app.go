package app

import "github.com/gin-gonic/gin"

var router *gin.Engine

func init() {
	router = gin.Default()
}

func StartApp() {
	UrlMappings()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
