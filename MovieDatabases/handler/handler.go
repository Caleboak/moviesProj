package handler

import (
	"MovieDatabases/entities"
	"MovieDatabases/repository"
	"MovieDatabases/services"

	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
)

type IMovieService interface {
	Create(ent entities.Movie) error
	ReadAll() (entities.DbMovie, error)
	ReadbyId(id string) (entities.Movie, error)
	DeletebyId(id string) error
	UpdatebyId(id string, ent entities.Movie) error
}

type MovieHandler struct {
	movieService IMovieService
}

func NewMovieHandler(mov services.MovieService) MovieHandler {
	return MovieHandler{
		movieService: &mov,
	}
}

func (m *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	ent := entities.Movie{}
	err := json.NewDecoder(r.Body).Decode(&ent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	//buf, _ := json.Marshal(&ent)
	//ioutil.WriteFile("moviedb.json", buf, 0644)

	err = m.movieService.Create(ent)
	if err != nil {
		switch err {
		case services.BadRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return

		case repository.NotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return

		case repository.ServerError:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("The movie has been created"))

}

func (m *MovieHandler) ReadAllMovie(w http.ResponseWriter, r *http.Request) {

	movie, err := m.movieService.ReadAll()

	if err != nil {
		switch err {
		case services.BadRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)

		case repository.NotFound:
			http.Error(w, err.Error(), http.StatusNotFound)

		case repository.ServerError:
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	Marshaled, _ := json.MarshalIndent(movie, "", " ")

	w.Write([]byte(Marshaled))
}

func (m *MovieHandler) ReadMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	movie, err := m.movieService.ReadbyId(id)

	if err != nil {
		switch err {
		case services.BadRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)

		case repository.NotFound:
			http.Error(w, err.Error(), http.StatusNotFound)

		case repository.ServerError:
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	Marshaled, _ := json.MarshalIndent(movie, "", " ")

	w.Write([]byte(Marshaled))
}

func (m *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := m.movieService.DeletebyId(id)

	if err != nil {
		switch err {
		case services.BadRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)

		case repository.NotFound:
			http.Error(w, err.Error(), http.StatusNotFound)

		case repository.ServerError:
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("The movie has been deleted"))

}

func (m *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ent := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&ent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = m.movieService.UpdatebyId(id, ent)
	if err != nil {
		switch err {
		case services.BadRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)

		case repository.NotFound:
			http.Error(w, err.Error(), http.StatusNotFound)

		case repository.ServerError:
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("The movie has been updated"))

}
