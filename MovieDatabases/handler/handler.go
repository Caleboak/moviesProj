package handler

import (
	"MovieDatabases/entities"
	"MovieDatabases/repo"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	ent := entities.Movie{}
	err := json.NewDecoder(r.Body).Decode(&ent)
	if err != nil {
		fmt.Println(err)
	}
	ent.SetId()
	jsonByte, err := repo.Create(ent)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)

}
