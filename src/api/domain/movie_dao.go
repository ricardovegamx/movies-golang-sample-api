package domain

import (
	"fmt"
)

var MovieDao MovieDaoInterface

type MovieDaoInterface interface {
	GetAll() *[]Movie
	Get(id uint64) (*Movie, error)
	Update(id uint64, movie *Movie) (*Movie, error)
	Create(movie *Movie) *Movie
}

type movieDao struct{}

func init() {
	MovieDao = &movieDao{}
}

func (m *movieDao) GetAll() *[]Movie {
	return &Database
}

func (m *movieDao) Get(id uint64) (*Movie, error) {
	for index, movie := range Database {
		if movie.Id == id {
			return &Database[index], nil
		}
	}

	return nil, NotFoundError(fmt.Sprintf("movie with id %d was not found", id))
}

func (m *movieDao) Update(id uint64, updatedMovie *Movie) (*Movie, error) {
	for index, movie := range Database {
		if movie.Id == id {
			Database[index] = *updatedMovie
			return &Database[index], nil
		}
	}

	return nil, BadRequestError(fmt.Sprintf("there was an error trying to update the movie with id %d", id))
}

func (m *movieDao) Create(movie *Movie) *Movie {
	Database = append(Database, *movie)
	return movie
}
