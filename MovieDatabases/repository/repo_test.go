package repository

import (
	"MovieDatabases/entities"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	file := "testFile.json"

	os.Create(file)
	log.Println("TestFile Created, Test started")

	exitVal := m.Run()
	log.Println("Tests ended, deleting TestFile...")
	os.Remove(file)
	os.Exit(exitVal)
}

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

	c := entities.DbMovie{}
	jsonMovie, _ := json.Marshal(&c)

	err := ioutil.WriteFile(file, jsonMovie, 0644)
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fatal()
	}

	r := MovieRepository{
		file,
	}

	//Act
	err = r.Create(jsonData1)
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fatal()
	}
	err = r.Create(jsonData2)
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fatal()
	}

	readFile, err := ioutil.ReadFile(file)
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fatal()
	}

	//Assert
	testStruct := entities.DbMovie{}
	json.Unmarshal(readFile, &testStruct)

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
		t.Fatal()
	}

	//Act
	dbEnt, err := r.GetAll()
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fatal()
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
		t.Fatal()
	}

	//Act
	entMovies1, err := r.GetById(id1)
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fatal()
	}
	entMovies2, err := r.GetById(id2)
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fatal()
	}

	//Assert
	assert.Equal(t, entMovies1, testdbStruct.Movies[0])
	assert.Equal(t, entMovies2, testdbStruct.Movies[1])
	assert.NotEqual(t, entMovies1, testdbStruct.Movies[1])
	assert.NotEqual(t, entMovies2, testdbStruct.Movies[0])

}

func TestDelete(t *testing.T) {
	//Arrange
	id := "200"

	file := "testFile.json"
	r := MovieRepository{
		file,
	}

	//Act
	err := r.Delete(id)
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fatal()
	}

	readFile, err := ioutil.ReadFile(file)
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fatal()
	}

	testdbStruct := entities.DbMovie{}

	err = json.Unmarshal(readFile, &testdbStruct)
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fatal()
	}

	//Assert
	assert.Equal(t, len(testdbStruct.Movies), 1)
}

func TestUpdate(t *testing.T) {
	//Arrange
	id := "100"

	updateStruct := entities.Movie{
		Id:          "100",
		Title:       "The 100",
		Genre:       []string{"Suspense", "Action"},
		Description: "The world becomes unsafe, they are forced to find livable surroundings",
		Director:    "Don't care",
		Actors:      []string{"Clarke", "Bellamy", "Hope"},
		Rating:      "89%",
	}

	file := "testFile.json"
	r := MovieRepository{
		file,
	}

	//Act
	err := r.Update(id, updateStruct)
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fatal()
	}

	readFile, err := ioutil.ReadFile(file)
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fatal()
	}

	testdbStruct := entities.DbMovie{}
	err = json.Unmarshal(readFile, &testdbStruct)
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fatal()
	}

	//Assert
	assert.Equal(t, updateStruct, testdbStruct.Movies[0])

}
