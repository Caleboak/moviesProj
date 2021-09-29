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
	if len(ent.Title) == 0 {
		return errors.New("bad request")
	} else if len(ent.Director) == 0 {
		return errors.New("bad request")
	}
	ent.SetId()
	return m.movieRepository.Create(ent)

}

func (m *MovieService) Read(id string) (entities.Movie, error) {
	errReturn := entities.Movie{}
	if len(id) == 0 {
		return errReturn, errors.New("bad request")
	}
	return m.movieRepository.FindById(id)
}
