package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HenrySaldanha/Go.FlashCards/models"
	"github.com/HenrySaldanha/Go.FlashCards/repository"
	"github.com/gorilla/mux"
)

// GetAllCards godoc
//
//	@Summary      List cards
//	@Description  get cards
//	@Tags         cards
//	@Accept       json
//	@Produce      json
//	@Success      200  {array}   model.Card
//	@Failure      400  {object}  httputil.HTTPError
//	@Failure      404  {object}  httputil.HTTPError
//	@Failure      500  {object}  httputil.HTTPError
//	@Router       /cards [get]
func GetAllCards(w http.ResponseWriter, r *http.Request) {
	cards := repository.GetAll()
	json.NewEncoder(w).Encode(cards)
}

// GetCardById godoc
//
//	@Summary      Get card by Id
//	@Description  Get card
//	@Tags         cards
//	@Accept       json
//	@Produce      json
//	@Param        id    path     string  true  "Card Id"
//	@Success      200  {object}   model.Card
//	@Failure      400  {object}  httputil.HTTPError
//	@Failure      404  {object}  httputil.HTTPError
//	@Failure      500  {object}  httputil.HTTPError
//	@Router       /cards/{id} [get]
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

// InsertCard godoc
//
//	@Summary      Insert a new cards
//	@Description  Insert cards
//	@Tags         cards
//	@Accept       json
//	@Produce      json
//	@Param        cards    body     model.Card  true  "Card Model"
//	@Success      200  {object}   model.Card
//	@Failure      400  {object}  httputil.HTTPError
//	@Failure      404  {object}  httputil.HTTPError
//	@Failure      500  {object}  httputil.HTTPError
//	@Router       /cards [post]
func InsertCard(w http.ResponseWriter, r *http.Request) {
	var card models.Card
	json.NewDecoder(r.Body).Decode(&card)
	card = repository.Save(card)
	json.NewEncoder(w).Encode(card)
}

// DeleteCard godoc
//
//	@Summary      Delete card by Id
//	@Description  Delete card
//	@Tags         cards
//	@Accept       json
//	@Produce      json
//	@Param        id    path     string  true  "Card Id"
//	@Success      202  {object}  model.Card
//	@Failure      400  {object}  httputil.HTTPError
//	@Failure      404  {object}  httputil.HTTPError
//	@Failure      500  {object}  httputil.HTTPError
//	@Router       /cards/{id} [delete]
func DeleteCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	repository.Delete(id)
	w.WriteHeader(http.StatusAccepted)
}

// UpdateCard godoc
//
//	@Summary      Update cards
//	@Description  Update cards
//	@Tags         cards
//	@Accept       json
//	@Produce      json
//	@Param        cards    body     model.Card  true  "Card Model"
//	@Success      200  {object}   model.Card
//	@Failure      400  {object}  httputil.HTTPError
//	@Failure      404  {object}  httputil.HTTPError
//	@Failure      500  {object}  httputil.HTTPError
//	@Router       /cards [put]
func UpdateCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var card models.Card
	json.NewDecoder(r.Body).Decode(&card)
	card = repository.Update(id, card)
	json.NewEncoder(w).Encode(card)
}
