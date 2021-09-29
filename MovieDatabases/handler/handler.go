package handler

import (
	"MovieDatabases/entities"
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

	err = m.movieService.Create(ent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func (m *MovieHandler) ReadMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	_, err := m.movieService.Read(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
