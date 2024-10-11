package data

import (
	"encoding/json"
	"io"
	"os"
	"wallps/utils"
)

const DATA_FILE_NAME string = "data.json"

type Collection struct {
	Nome         string
	Tipo         string   // pode ser simples ou explicita
	Path         string   // só vale para o tipo de coleção simples
	Images       []string // só vale para o tipo coleção explicita
	CurrentImage int      // vale tanto para coleção simples quanto explícita
	ChangePeriod int
}

type ProgramData struct {
	Collections       []Collection
	CurrentCollection int
	AutoSwitch        bool
}

// returns the stored program data
func Data() ProgramData {
	dataFile := OpenDataFile()
	data := utils.Expect(io.ReadAll(dataFile))
	var dataBuffer ProgramData
	err := json.Unmarshal(data, &dataBuffer)
	if err != nil {
		panic(err)
	}
	return dataBuffer
}

// returns the data file descriptor for reading only
func OpenDataFile() *os.File {

	pd := utils.ProgramDirectory()        //the full path of the data file directory
	dataFile := pd + "/" + DATA_FILE_NAME //the full path of the data file

	file, err := os.Open(dataFile)
	if err == nil {
		return file
	}

	//if the data file does not exist yet:
	file = utils.Expect(os.Create(dataFile))
	EmptyData := utils.Expect(json.Marshal(ProgramData{}))
	utils.Expect(file.Write(EmptyData))
	file.Close()
	return OpenDataFile()
}

func WriteDataFile(data ProgramData) {
	pd := utils.ProgramDirectory()
	dataFile := pd + "/" + DATA_FILE_NAME
	file := utils.Expect(os.Create(dataFile))
	byteData := utils.Expect(json.Marshal(data))
	file.Write(byteData)
	file.Close()
}

// change the collections's current image and returns it
func (col *Collection) NextImage() int {
	collectionDir := utils.Expect(os.ReadDir(col.Path))
	col.CurrentImage++
	if col.CurrentImage >= len(collectionDir) {
		col.CurrentImage = 0
	}
	return col.CurrentImage
}
