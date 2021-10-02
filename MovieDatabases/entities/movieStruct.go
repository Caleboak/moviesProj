package entities

import "github.com/google/uuid"

type Movie struct {
	Id          string
	Title       string   `validate:"required"`
	Genre       []string `validate:"required"`
	Description string   `validate:"required"`
	Director    string   `validate:"required"`
	Actors      []string `validate:"required"`
	Rating      string   `validate:"required"`
}

func (mov *Movie) SetId() {

	mov.Id = uuid.New().String()
}
