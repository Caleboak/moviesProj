package repo

import (
	"MovieDatabases/entities"
	"MovieDatabases/fileConverter"
	"io/ioutil"
)

func Create(userMovie entities.Movie) ([]byte, error) {
	var errReturn []byte
	fc := fileConverter.NewFileConverter("moviedb.json")

	dbStruct, err := fc.ConvertToGo()
	if err != nil {
		return errReturn, err
	}
	dbStruct.Movies = append(dbStruct.Movies, userMovie)
	jsonByte, err := fc.ConvertToFile(dbStruct)
	if err != nil {
		return errReturn, err
	}

	ioutil.WriteFile("moviedb.json", jsonByte, 0644)

	return jsonByte, nil

}
