package repository

import (
	"fmt"

	"github.com/HenrySaldanha/Go.FlashCards/database"
	"github.com/HenrySaldanha/Go.FlashCards/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAll() []model.Card {
	fmt.Println("Getting all cards")

	var cards []model.Card

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

func GetById(id string) (card *model.Card) {
	fmt.Printf("Getting by id %s\n", id)

	var cards []model.Card

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

func Save(card model.Card) model.Card {
	fmt.Printf("Saving a new card %+v\n", card)

	card.Id = uuid.New().String()

	_, err := database.InsertOne(database.CollectionName, card)

	if err != nil {
		panic(err)
	}

	return card
}

func Delete(id string) {
	fmt.Printf("Deleting a card %s\n", id)

	filter := bson.D{{Key: "id", Value: id}}
	err := database.DeleteOne(database.CollectionName, filter)

	if err != nil {
		panic(err)
	}
}

func Update(id string, card model.Card) model.Card {
	fmt.Printf("Updating a card %s\n", id)

	filter := bson.D{{Key: "id", Value: id}}
	card.Id = id
	_, err := database.ReplaceDocument(database.CollectionName, filter, card)

	if err != nil {
		panic(err)
	}

	return card
}
