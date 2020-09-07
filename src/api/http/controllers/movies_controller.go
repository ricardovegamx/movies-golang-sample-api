package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ricardovegamx/movies-api/src/api/services"
)

/*
Return list of all movies
*/
func Index(c *gin.Context) {
	movies := services.MoviesService.GetAll()

	c.JSON(http.StatusOK, gin.H{"data": movies})

	return
}
