package services

import (
	"MovieDatabases/entities"
	"MovieDatabases/repository"

	"gopkg.in/go-playground/validator.v9"
)

type IMovieRepository interface {
	Create(userMovie entities.Movie) error
	GetAll() (entities.DbMovie, error)
	GetById(passedId string) (entities.Movie, error)
	Delete(passedId string) error
	Update(id string, ent entities.Movie) error
}

type MovieService struct {
	movieRepository IMovieRepository
}

func NewMovieServices(repo repository.MovieRepository) MovieService {
	return MovieService{
		movieRepository: &repo,
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

	return m.movieRepository.GetAll()
}

func (m *MovieService) ReadbyId(id string) (entities.Movie, error) {
	errReturn := entities.Movie{}
	if len(id) == 0 {
		return errReturn, BadRequest
	}
	return m.movieRepository.GetById(id)
}

func (m *MovieService) DeletebyId(id string) error {

	if len(id) == 0 {
		return BadRequest
	}
	return m.movieRepository.Delete(id)
}

func (m *MovieService) UpdatebyId(id string, ent entities.Movie) error {
	if len(id) == 0 {
		return BadRequest
	}
	validate := validator.New()
	err := validate.Struct(ent)
	if err != nil {
		return BadRequest
	}

	return m.movieRepository.Update(id, ent)

}
