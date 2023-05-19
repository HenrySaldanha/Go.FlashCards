package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HenrySaldanha/Go.FlashCards/models"
	"github.com/HenrySaldanha/Go.FlashCards/repository"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllCards(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Buscando todos os registros")
	repository.GetAllCards()
}

func InsertCard(w http.ResponseWriter, r *http.Request) {
	var card models.Card
	json.NewDecoder(r.Body).Decode(&card)

	repository.SaveCard(card)
}
