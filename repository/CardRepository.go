package repository

import (
	"fmt"

	"github.com/HenrySaldanha/Go.FlashCards/database"
	"github.com/HenrySaldanha/Go.FlashCards/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCards() ([]models.Card, error) {
	fmt.Println("getting all cards")

	client, ctx, cancel, err := database.GetMongoDBConnection()
	if err != nil {
		panic(err)
	}

	defer database.Close(client, ctx, cancel)

	var filter, option interface{}
	filter = bson.D{}
	option = bson.D{{"_id", 0}}

	cursor, err := database.Query(client, ctx, database.CollectionName, filter, option)
	if err != nil {
		panic(err)
	}

	var results []bson.D

	if err := cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	fmt.Println("Query Result")
	for _, doc := range results {
		fmt.Println(doc)
	}

	return nil, nil
}

func SaveCard(card models.Card) {
	fmt.Printf("Saving a new card %+v\n", card)

	card.Id = uuid.New().String()

	client, ctx, cancel, err := database.GetMongoDBConnection()
	if err != nil {
		panic(err)
	}

	defer database.Close(client, ctx, cancel)

	insertOneResult, err := database.InsertOne(client, ctx, database.CollectionName, card)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Result of InsertOne %s\n", insertOneResult.InsertedID)
}
