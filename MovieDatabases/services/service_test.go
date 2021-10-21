package services

import (
	"MovieDatabases/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	jsonData1 := entities.Movie{
		Id:          "",
		Title:       "",
		Genre:       []string{"suspense"},
		Description: "People are forced to play a game to win money",
		Director:    "Don't know",
		Actors:      []string{"Sang Woo"},
		Rating:      "78%",
	}

	s := MovieService{}

	err := s.Create(jsonData1)
	assert.EqualErrorf(t, err, BadRequest.Error(), "Error should be %v, but got %v", BadRequest.Error(), err)

	//Need to test for the setId also

}

func TestReadById(t *testing.T) {
	//Arrange
	id := ""

	//Act
	s := MovieService{}
	_, err := s.ReadbyId(id)

	//Assert
	assert.EqualErrorf(t, err, BadRequest.Error(), "Error should return %v, but got %v", BadRequest.Error(), err)

}
