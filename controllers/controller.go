package controllers

import (
	"net/http"

	"github.com/HenrySaldanha/Go.FlashCards/model"
	"github.com/HenrySaldanha/Go.FlashCards/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
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
func GetAllCards(c *gin.Context) {
	cards := repository.GetAll()

	c.JSON(http.StatusOK, cards)
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
func GetCardById(c *gin.Context) {
	id := c.Params.ByName("id")

	card := repository.GetById(id)
	if card != nil {
		c.JSON(http.StatusOK, card)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Card is Not found"})
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
func InsertCard(c *gin.Context) {
	var card model.Card

	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	card = repository.Save(card)
	c.JSON(http.StatusOK, card)
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
func DeleteCard(c *gin.Context) {
	id := c.Params.ByName("id")
	repository.Delete(id)
	c.JSON(http.StatusAccepted, gin.H{"data": "Card has been deleted"})
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
func UpdateCard(c *gin.Context) {
	id := c.Params.ByName("id")
	var card model.Card

	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	card = repository.Update(id, card)
	c.JSON(http.StatusOK, card)
}
