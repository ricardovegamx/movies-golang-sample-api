package domain

var MovieDao MovieDaoInterface

type MovieDaoInterface interface {
	GetAll() *[]Movie
}

type movieDao struct{}

func init() {
	MovieDao = &movieDao{}
}

func (movie *movieDao) GetAll() *[]Movie {
	return &Database
}
