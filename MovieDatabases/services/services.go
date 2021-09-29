package services

import (
	"MovieDatabases/entities"
	"MovieDatabases/repository"
	"errors"
)

type MovieService struct {
	movieRepository repository.MovieRepository
}

func NewMovieServices(repo repository.MovieRepository) MovieService {
	return MovieService{
		movieRepository: repo,
	}
}

func (m *MovieService) Create(ent entities.Movie) error {
	ent.SetId()
	return errors.New("")

}
