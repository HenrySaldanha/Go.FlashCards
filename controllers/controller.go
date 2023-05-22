package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HenrySaldanha/Go.FlashCards/models"
	"github.com/HenrySaldanha/Go.FlashCards/repository"
	"github.com/gorilla/mux"
)

func GetAllCards(w http.ResponseWriter, r *http.Request) {
	cards := repository.GetAll()
	json.NewEncoder(w).Encode(cards)
}

func GetCardById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	card := repository.GetById(id)
	if card != nil {
		json.NewEncoder(w).Encode(&card)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func InsertCard(w http.ResponseWriter, r *http.Request) {
	var card models.Card
	json.NewDecoder(r.Body).Decode(&card)
	card = repository.Save(card)
	json.NewEncoder(w).Encode(card)
}

func DeleteCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	repository.Delete(id)
	w.WriteHeader(http.StatusAccepted)
}
