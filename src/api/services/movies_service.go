package services

import (
	"github.com/ricardovegamx/movies-golang-sample-api/src/api/domain"
)

type moviesService struct{}

var MoviesService moviesService

func (m *moviesService) GetAll() *[]domain.Movie {
	return domain.MovieDao.GetAll()
}

func (m *moviesService) Get(id uint64) (*domain.Movie, error) {
	movie, err := domain.MovieDao.Get(id)

	if err != nil {
		return nil, err
	}

	return movie, err
}

func (m *moviesService) Update(id uint64, movie *domain.Movie) (*domain.Movie, error) {
	updatedMovie, err := domain.MovieDao.Update(id, movie)

	if err != nil {
		return nil, err
	}

	return updatedMovie, err
}

func (m *moviesService) Create(movie *domain.Movie) *domain.Movie {
	return domain.MovieDao.Create(movie)
}
