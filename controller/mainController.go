package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"../service"
	"../model"
)

func showAllLiterature(w http.ResponseWriter, r *http.Request) {

	literatureSlice := service.ReceiveAllContent()

	json.NewEncoder(w).Encode(literatureSlice)

}

func showLiteratureByName(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var literature, requestedLiterature model.Literature
	var response model.Response

	err := decoder.Decode(&literature)
	if err != nil {
		log.Println("Decoding JSON error.")

		response.Success = false
		response.Message = "Data was not saved."
	} else {
		requestedLiterature, response = service.ReceiveLitByName(literature)
	}

	if response.Success {
		json.NewEncoder(w).Encode(requestedLiterature)
	} else {
		json.NewEncoder(w).Encode(response)
	}

}

func showLiteratureByType(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var literature model.Literature
	var literatureSlice []model.Literature
	var response model.Response

	err := decoder.Decode(&literature)
	if err != nil {
		log.Println("Decoding JSON error.")

		response.Success = false
		response.Message = "Data was not saved."
	} else {
		literatureSlice = service.ReceiveLitByType(literature)
		response.Success = true
	}

	if response.Success {
		json.NewEncoder(w).Encode(literatureSlice)
	} else {
		json.NewEncoder(w).Encode(response)
	}

}

func storeLiterature(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var literature model.Literature
	var response model.Response

	err := decoder.Decode(&literature)
	if err != nil {
		log.Println("Decoding JSON error.")

		response.Success = false
		response.Message = "Data was not saved."
	} else {
		response = service.CreateNewItem(literature)
	}

	json.NewEncoder(w).Encode(response)

}

func MainController() {

	http.HandleFunc("/literature", showAllLiterature)
	http.HandleFunc("/store", storeLiterature)
	http.HandleFunc("/getlitname", showLiteratureByName)
	http.HandleFunc("/getlittype", showLiteratureByType)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
