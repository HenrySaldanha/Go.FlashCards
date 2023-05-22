package repository

import (
	"fmt"

	"github.com/HenrySaldanha/Go.FlashCards/database"
	"github.com/HenrySaldanha/Go.FlashCards/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAll() []models.Card {
	fmt.Println("getting all cards")

	var cards []models.Card

	var filter, option interface{}
	filter = bson.D{}
	option = bson.D{{Key: "_id", Value: 0}}

	cursor, ctx, err := database.Query(database.CollectionName, filter, option)
	if err != nil {
		panic(err)
	}

	if err := cursor.All(ctx, &cards); err != nil {
		panic(err)
	}

	return cards
}

func GetById(id string) (card *models.Card) {
	fmt.Println("getting all cards")

	var cards []models.Card

	var filter, option interface{}
	filter = bson.D{{Key: "id", Value: id}}
	option = bson.D{{Key: "_id", Value: 0}}

	cursor, ctx, err := database.Query(database.CollectionName, filter, option)
	if err != nil {
		panic(err)
	}

	if err := cursor.All(ctx, &cards); err != nil {
		panic(err)
	}

	if len(cards) > 0 {
		return &cards[0]
	}

	return nil
}

func Save(card models.Card) models.Card {
	fmt.Printf("Saving a new card %+v\n", card)

	card.Id = uuid.New().String()

	_, err := database.InsertOne(database.CollectionName, card)

	if err != nil {
		panic(err)
	}

	return card
}
