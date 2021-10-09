package repository

import (
	"MovieDatabases/entities"
	"encoding/json"
	"io/ioutil"
)

type IMovieRepository interface {
	Create(userMovie entities.Movie) error
	GetAll() (entities.DbMovie, error)
	GetById(passedId string) (entities.Movie, error)
	Delete(passedId string) error
	Update(id string, ent entities.Movie) error
}

type MovieRepository struct {
	filename string
}

func NewMovieRepository(fn string) IMovieRepository {
	return &MovieRepository{
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

func (r *MovieRepository) GetAll() (entities.DbMovie, error) {
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

func (r *MovieRepository) GetById(passedId string) (entities.Movie, error) {
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

func (r *MovieRepository) Delete(passedId string) error {
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

func (r *MovieRepository) Update(id string, ent entities.Movie) error {
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}
	dbStruct := entities.DbMovie{}

	err = json.Unmarshal(file, &dbStruct)
	if err != nil {
		return ServerError
	}

	for i, v := range dbStruct.Movies {
		if v.Id == id {
			v = ent
			ent.Id = id
			entStruct := entities.DbMovie{}                                     //creating a entStruct to enable us append ent as a slice struct
			entStruct.Movies = append(entStruct.Movies, ent)                    //add the ent to its struct slice
			dbStruct.Movies = append(dbStruct.Movies[:i], entStruct.Movies...)  //append with the previous
			dbStruct.Movies = append(dbStruct.Movies, dbStruct.Movies[i+1:]...) //append with the one after
			Marshaled, err := json.MarshalIndent(dbStruct, "", " ")
			if err != nil {
				return ServerError
			}
			ioutil.WriteFile(r.filename, Marshaled, 0644)
			return nil
		}
	}
	return NotFound
}
