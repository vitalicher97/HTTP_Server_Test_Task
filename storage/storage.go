package storage

import (
	"../model"
	"log"
)

var litMemory map[string]string

func InitLitMemory() {
	litMemory = make(map[string]string)

	log.Println("In-memory storage was initialised.")

	/*	litMemory["50 shades of Go"] = "ARTICLE"
		litMemory["Learning Go"] = "BOOK"*/

}

func StoreLiterature(literature model.Literature) model.Response {
	litMemory[literature.Name] = literature.LitType

	var response model.Response
	response.Success = true
	response.Message = "Data was saved."

	log.Println("Data saved in in-memory storage.")

	return response
}

func GetAllLiterature() []model.Literature {
	var literature model.Literature
	var literatureSlice []model.Literature
	for key, value := range litMemory {
		literature.Name = key
		literature.LitType = value
		literatureSlice = append(literatureSlice, literature)
	}
	return literatureSlice
}

func GetLiteratureByName(name string) (model.Literature, model.Response) {

	var literature model.Literature
	var response model.Response

	value, ok := litMemory[name]
	if ok {
		literature.Name = name
		literature.LitType = value

		response.Success = true
		response.Message = "Record was found."
		return literature, response
	} else {
		response.Success = false
		response.Message = "There is no record with provided name."
	}

	return literature, response

}

func GetLiteratureByType(litType string) []model.Literature {

	var literature model.Literature
	var literatureSlice []model.Literature

	for key, value := range litMemory {
		if litType == value {
			literature.Name = key
			literature.LitType = value
			literatureSlice = append(literatureSlice, literature)
		}
	}

	return literatureSlice

}

func IsNameExist(name string) bool {
	_, ok := litMemory[name]
	return ok
}
