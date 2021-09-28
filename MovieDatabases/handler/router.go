package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer() *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/POST", CreateMovie).Methods("POST")

	svr := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
	}

	return svr

}
