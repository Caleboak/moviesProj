package repository

import (
	"MovieDatabases/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {

	//Arrange
	jsonData1 := entities.Movie{
		Id:          "100",
		Title:       "Squid Game",
		Genre:       []string{"suspense"},
		Description: "People are forced to play a game to win money",
		Director:    "Don't know",
		Actors:      []string{"Sang Woo"},
		Rating:      "78%",
	}

	jsonData2 := entities.Movie{
		Id:          "200",
		Title:       "On my Block",
		Genre:       []string{"suspense"},
		Description: "Bunch of kids looking for a treasure",
		Director:    "Blah Blah",
		Actors:      []string{"Monsee"},
		Rating:      "75%",
	}
	file := "testFile.json"

	os.Create(file)

	c := entities.DbMovie{}
	jsonMovie, _ := json.Marshal(&c)

	err := ioutil.WriteFile(file, jsonMovie, 0644)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	r := MovieRepository{
		file,
	}

	//Act
	err = r.Create(jsonData1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	err = r.Create(jsonData2)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	readFile, err := ioutil.ReadFile(file)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	//Assert
	testStruct := entities.DbMovie{}
	json.Unmarshal(readFile, &testStruct)
	fmt.Println(testStruct)
	assert.Equal(t, jsonData1, testStruct.Movies[0])
	assert.Equal(t, jsonData2, testStruct.Movies[1])

}

func TestGetAll(t *testing.T) {

	//Arrange
	file := "testFile.json"
	r := MovieRepository{
		file,
	}

	testdbStruct := entities.DbMovie{}
	fileData, _ := ioutil.ReadFile(file)
	err := json.Unmarshal(fileData, &testdbStruct)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	//Act
	dbEnt, err := r.GetAll()
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	//Assert
	assert.Equal(t, testdbStruct.Movies[0], dbEnt.Movies[0])
	assert.Equal(t, testdbStruct.Movies[1], dbEnt.Movies[1])
	assert.NotEqual(t, testdbStruct.Movies[0], dbEnt.Movies[1])
	assert.NotEqual(t, testdbStruct.Movies[1], dbEnt.Movies[0])

}

func TestGetById(t *testing.T) {
	//Arrange
	id1 := "100"
	id2 := "200"

	file := "testFile.json"
	r := MovieRepository{
		file,
	}

	testdbStruct := entities.DbMovie{}
	fileData, _ := ioutil.ReadFile(file)
	err := json.Unmarshal(fileData, &testdbStruct)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	entMovies1, err := r.GetById(id1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	entMovies2, err := r.GetById(id2)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	//Assert
	assert.Equal(t, entMovies1, testdbStruct.Movies[0])
	assert.Equal(t, entMovies2, testdbStruct.Movies[1])
	assert.NotEqual(t, entMovies1, testdbStruct.Movies[1])
	assert.NotEqual(t, entMovies2, testdbStruct.Movies[0])

}
