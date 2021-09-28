package fileConverter

import (
	"MovieDatabases/entities"
	"encoding/json"
	"io/ioutil"
)

type FileConverter struct {
	filename string
}

func NewFileConverter(fn string) FileConverter {
	return FileConverter{
		filename: fn,
	}
}

func (f *FileConverter) ConvertToGo() (entities.DbMovie, error) {
	dbEnt := entities.DbMovie{}
	file, err := ioutil.ReadFile(f.filename)
	if err != nil {
		return dbEnt, err
	}

	err = json.Unmarshal(file, &dbEnt)
	if err != nil {
		return dbEnt, err
	}
	return dbEnt, nil
}

func (f *FileConverter) ConvertToFile(ent entities.DbMovie) ([]byte, error) {
	var errReturn []byte
	Marshaled, err := json.Marshal(&ent)
	if err != nil {
		return errReturn, err
	}

	return Marshaled, err

}
