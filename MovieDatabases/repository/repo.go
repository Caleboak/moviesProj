package repository

import (
	"MovieDatabases/entities"
	"encoding/json"
	"errors"
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
		return err
	}

	err = json.Unmarshal(file, &dbEnt)
	if err != nil {
		return err
	}

	dbEnt.Movies = append(dbEnt.Movies, userMovie)

	Marshaled, err := json.Marshal(&dbEnt)
	if err != nil {
		return err
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil

}

func (r *MovieRepository) FindById(passedId string) (entities.Movie, error) {
	dbEnt := entities.DbMovie{}
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return dbEnt.Movies[0], err
	}

	err = json.Unmarshal(file, &dbEnt)
	if err != nil {
		return dbEnt.Movies[0], err
	}

	for _, v := range dbEnt.Movies {
		if v.Id == passedId {
			return v, nil
		}
	}

	return dbEnt.Movies[0], errors.New("id not found")
}
