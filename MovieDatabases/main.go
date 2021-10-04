package main

import (
	"MovieDatabases/handler"
	"MovieDatabases/repository"
	"MovieDatabases/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	filename := "moviedb.json"
	repo := repository.NewMovieRepository(filename)
	serv := services.NewMovieServices(repo)
	handle := handler.NewMovieHandler(serv)
	r := mux.NewRouter()
	r.HandleFunc("/Movie", handle.CreateMovie).Methods("POST")
	r.HandleFunc("/Movie", handle.ReadAllMovie).Methods("GET")
	r.HandleFunc("/Movie/{id}", handle.ReadMovie).Methods("GET")
	r.HandleFunc("/Movie/{id}", handle.DeleteMovie).Methods("DELETE")
	svr := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
	}
	svr.ListenAndServe()

}
