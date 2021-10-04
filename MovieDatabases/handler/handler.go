package handler

import (
	"MovieDatabases/entities"
	"MovieDatabases/repository"
	"MovieDatabases/services"
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
)

type MovieHandler struct {
	movieService services.MovieService
}

func NewMovieHandler(mov services.MovieService) MovieHandler {
	return MovieHandler{
		movieService: mov,
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

		case repository.NotFound:
			http.Error(w, err.Error(), http.StatusNotFound)

		case repository.ServerError:
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
