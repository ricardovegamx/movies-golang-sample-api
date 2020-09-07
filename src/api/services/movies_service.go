package services

import (
	"github.com/ricardovegamx/movies-golang-sample-api/src/api/domain"
)

type moviesService struct{}

var MoviesService moviesService

func (m *moviesService) GetAll() *[]domain.Movie {
	return domain.MovieDao.GetAll()
}
