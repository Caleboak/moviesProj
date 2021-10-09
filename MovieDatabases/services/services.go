package services

import (
	"MovieDatabases/entities"
	"MovieDatabases/repository"

	"gopkg.in/go-playground/validator.v9"
)

type IMovieService interface {
	Create(ent entities.Movie) error
	ReadAll() (entities.DbMovie, error)
	ReadbyId(id string) (entities.Movie, error)
	DeletebyId(id string) error
	UpdatebyId(id string, ent entities.Movie) error
}

type MovieService struct {
	movieRepository repository.IMovieRepository
}

func NewMovieServices(repo repository.IMovieRepository) IMovieService {
	return &MovieService{
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
