package dao

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/St0iK/go-quote-parser/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_DB_URL")))
// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// err = client.Connect(ctx)
// collection := client.Database("testing").Collection("quotes")

var collection *mongo.Collection

const (
	// DBNAME ...
	DBNAME = "quotes-parser"
	// COLLECTION ...
	COLLECTION = "quotes"
)

// Connect ...
func Connect() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_DB_URL")))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(DBNAME).Collection(COLLECTION)
	log.Println("Connected to MongoDB!")

}


func Insert(quote model.Quote) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := collection.InsertOne(ctx, bson.M{
		"author": quote.Author,
		"quote":  quote.QuoteText,
	})

	return err
}
// Insert a movie into database
func InsertV1(quote model.QuoteV1) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{
		"author": quote.Author,
		"quote": quote.Quote,
	})

	id := res.InsertedID
	fmt.Println(id)
	return err
}

// Insert a movie into database
func InsertV2(quote model.QuoteV2) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{
		"author": quote.Author,
		"quote": quote.Quote,
	})

	id := res.InsertedID
	fmt.Println(id)
	return err
}
