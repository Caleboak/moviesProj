package entities

import "github.com/google/uuid"

type Movie struct {
	Id          string
	Title       string
	Genre       []string
	Description string
	Director    string
	Actors      []string
	Rating      string
}

func (mov *Movie) SetId() {

	mov.Id = uuid.New().String()
}
