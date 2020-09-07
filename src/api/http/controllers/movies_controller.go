package controllers

import "C"
import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ricardovegamx/movies-golang-sample-api/src/api/domain"
	"github.com/ricardovegamx/movies-golang-sample-api/src/api/http/responses"
	"github.com/ricardovegamx/movies-golang-sample-api/src/api/services"
)

/*
Return list of all movies
*/
func Index(c *gin.Context) {
	movies := services.MoviesService.GetAll()

	c.JSON(http.StatusOK, &responses.SuccessResponseGetAll{
		Status: responses.StatusSuccess,
		Data: struct {
			Movies *[]domain.Movie `json:"movies"`
		}{Movies: movies},
	})

	return
}

func Get(c *gin.Context) {
	movieId, _ := strconv.ParseUint(c.Param("movie_id"), 10, 64)
	movie, err := services.MoviesService.Get(movieId)

	if err != nil {
		c.JSON(http.StatusNotFound, &responses.FailResponse{
			Status: responses.StatusFail,
			Data:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &responses.SuccessResponseGet{
		Status: responses.StatusSuccess,
		Data: struct {
			Movie *domain.Movie `json:"movie"`
		}{Movie: movie},
	})
	return
}

func Update(c *gin.Context) {
	movieId, _ := strconv.ParseUint(c.Param("movie_id"), 10, 64)
	movie := domain.Movie{}
	if c.ShouldBind(&movie) == nil {
		updatedMovie, err := services.MoviesService.Update(movieId, &movie)
		if err != nil {
			c.JSON(http.StatusBadRequest, &responses.FailResponse{
				Status: responses.StatusFail,
				Data:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &responses.SuccessResponseGet{
			Status: responses.StatusSuccess,
			Data: struct {
				Movie *domain.Movie `json:"movie"`
			}{Movie: updatedMovie},
		})
		return
	}

	c.JSON(http.StatusInternalServerError, &responses.ErrorResponse{
		Status:  responses.StatusError,
		Message: "Server did not respond",
	})
	return
}

func Create(c *gin.Context) {
	movie := &domain.Movie{}
	if c.ShouldBind(movie) == nil {
		createdMovie := services.MoviesService.Create(movie)
		c.JSON(http.StatusCreated, &responses.SuccessResponseGet{
			Status: responses.StatusSuccess,
			Data: struct {
				Movie *domain.Movie `json:"movie"`
			}{Movie: createdMovie},
		})
		return
	}

	c.JSON(http.StatusBadRequest, &responses.FailResponse{
		Status: responses.StatusFail,
		Data:   "malformed body",
	})
	return
}
