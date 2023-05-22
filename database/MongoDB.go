package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const DatabaseName = "FlashCards"
const CollectionName = "cards"
const MongoDbConnectionString = "mongodb://localhost:27017"

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func GetMongoDBConnection() (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDbConnectionString))
	return client, ctx, cancel, err
}

func Ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

func InsertOne(collectionName string, doc interface{}) (*mongo.InsertOneResult, error) {

	client, ctx, cancel, err := GetMongoDBConnection()
	if err != nil {
		panic(err)
	}

	defer Close(client, ctx, cancel)

	collection := client.Database(DatabaseName).Collection(collectionName)
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func Query(collectionName string, query, field interface{}) (result *mongo.Cursor, ctx context.Context, err error) {
	client, ctx, cancel, err := GetMongoDBConnection()
	if err != nil {
		panic(err)
	}

	defer Close(client, ctx, cancel)

	// select database and collection.
	collection := client.Database(DatabaseName).Collection(collectionName)

	// collection has an method Find,
	// that returns a mongo.cursor
	// based on query and field.
	result, err = collection.Find(ctx, query,
		options.Find().SetProjection(field))
	return result, ctx, err
}
