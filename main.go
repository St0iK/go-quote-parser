package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"os"
	"github.com/St0iK/go-quote-parser/model"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	collection := client.Database("testing").Collection("quotes")

	jsonFile, err := os.Open("quotes/quotations.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var quotes []model.Quote

	err = json.Unmarshal(byteValue, &quotes)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(quotes))

	for i := 0; i < len(quotes); i++ {
		ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
		res, _ := collection.InsertOne(ctx, bson.M{"name": quotes[i].Name, "text": quotes[i].Text})
		id := res.InsertedID
		fmt.Println(id)
		fmt.Println("\nName: " + quotes[i].Name)
		fmt.Println("\nText: " + quotes[i].Text)

	}

	// Set ENV Variables to HandleConnection to MongoDB
	// ToDo Move models to the right file
	// Move MongoDB Operations to their own file
	// Add parsers for the other files

}
