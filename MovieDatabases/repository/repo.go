package repository

import (
	"MovieDatabases/entities"
	"encoding/json"
	"io/ioutil"
)

type MovieRepository struct {
	filename string
}

func NewMovieRepository(fn string) MovieRepository {
	return MovieRepository{
		filename: fn,
	}
}

func (r *MovieRepository) Create(userMovie entities.Movie) error {

	dbEnt := entities.DbMovie{}
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}

	err = json.Unmarshal(file, &dbEnt)
	if err != nil {
		return ServerError
	}

	dbEnt.Movies = append(dbEnt.Movies, userMovie)

	Marshaled, err := json.MarshalIndent(&dbEnt, "", " ")
	if err != nil {
		return ServerError
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil
}

func (r *MovieRepository) FindAll() (entities.DbMovie, error) {
	dbEnt := entities.DbMovie{}
	errEnt := entities.DbMovie{} //returns incase of error
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return errEnt, ServerError
	}

	err = json.Unmarshal(file, &dbEnt)
	if err != nil {
		return errEnt, ServerError
	}

	return dbEnt, nil
}

func (r *MovieRepository) FindById(passedId string) (entities.Movie, error) {
	dbEnt := entities.DbMovie{}
	errEnt := entities.Movie{} //returns incase of error
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return errEnt, ServerError
	}

	err = json.Unmarshal(file, &dbEnt)
	if err != nil {
		return errEnt, ServerError
	}

	for _, v := range dbEnt.Movies {
		if v.Id == passedId {
			return v, nil
		}
	}

	return errEnt, NotFound
}

func (r *MovieRepository) DeleteById(passedId string) error {
	dbEnt := entities.DbMovie{}
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}

	err = json.Unmarshal(file, &dbEnt)
	if err != nil {
		return ServerError
	}

	for i, v := range dbEnt.Movies {
		if v.Id == passedId {
			dbEnt.Movies = append(dbEnt.Movies[:i], dbEnt.Movies[i+1:]...)
			Marshaled, err := json.MarshalIndent(&dbEnt, "", " ")
			if err != nil {
				return ServerError
			}
			ioutil.WriteFile(r.filename, Marshaled, 0644)
			return nil
		}
	}

	return NotFound
}
