package controllers

import (
	"fmt"
	"net/http"

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
	fmt.Fprint(w, "Inserindo registro")
}
