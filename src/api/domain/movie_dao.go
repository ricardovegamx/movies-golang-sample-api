package domain

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MovieDao MovieDaoInterface

type MovieDaoInterface interface {
	GetAll() *[]Movie
	Get(id uint64) (*Movie, error)
	Update(id uint64, movie *Movie) (*Movie, error)
	Create(movie *Movie) *Movie
	Delete(id uint64) bool
}

type movieDao struct{}

func init() {
	MovieDao = &movieDao{}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
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

func (m *movieDao) Delete(id uint64) bool {
	for index, movie := range Database {
		if movie.Id == id {
			Database = append(Database[:index], Database[index+1:]...)
			return true
		}
	}

	return false
}
