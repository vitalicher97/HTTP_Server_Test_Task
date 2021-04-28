package service

import (
	"../storage"
	"../model"
	"log"
)

func ReceiveAllContent() []model.Literature {
	allLiterature := storage.GetAllLiterature()
	return allLiterature
}

func ReceiveLitByName(literature model.Literature) (model.Literature, model.Response) {
	requestedLiterature, response := storage.GetLiteratureByName(literature.Name)
	return requestedLiterature, response
}

func ReceiveLitByType(literature model.Literature) []model.Literature {
	literatureSlice := storage.GetLiteratureByType(literature.LitType)
	return literatureSlice
}

func CreateNewItem(literature model.Literature) model.Response {

	var response model.Response

	if literature.LitType != "ARTICLE" && literature.LitType != "BOOK" {
		log.Println("Literature type is not allowed: " + literature.LitType)

		response.Success = false
		response.Message = "Data was not saved. Only types 'ARTICLE' or 'BOOK' are allowed."

		return response
	}

	if storage.IsNameExist(literature.Name) {
		log.Println("Provided literature name already exists. Duplicates are not allowed.")

		response.Success = false
		response.Message = "Data was not saved. Record with provided name already exists."

		return response
	}

	response = storage.StoreLiterature(literature)

	return response

}
