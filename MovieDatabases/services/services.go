package services

import (
	"MovieDatabases/entities"
	"MovieDatabases/repository"

	"gopkg.in/go-playground/validator.v9"
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

	validate := validator.New()
	err := validate.Struct(ent)
	if err != nil {
		return BadRequest
	}

	ent.SetId()
	return m.movieRepository.Create(ent)

}

func (m *MovieService) ReadAll() (entities.DbMovie, error) {

	return m.movieRepository.FindAll()
}

func (m *MovieService) ReadbyId(id string) (entities.Movie, error) {
	errReturn := entities.Movie{}
	if len(id) == 0 {
		return errReturn, BadRequest
	}
	return m.movieRepository.FindById(id)
}
