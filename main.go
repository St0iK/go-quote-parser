package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/St0iK/go-quote-parser/dao"
	"github.com/St0iK/go-quote-parser/model"
)

func init() {
	dao.Connect()
}

func main() {

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
		fmt.Printf("%+v\n", quotes[i])
		dao.Insert(quotes[i])
		// ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
		// res, _ = collection.InsertOne(ctx, bson.M{"name": quotes[i].Name, "text": quotes[i].Text})
		// id := res.InsertedID
		// fmt.Println(id)
		// dao.Insert(quotes[i])
		fmt.Println("\nName: " + quotes[i].Name)
		// fmt.Println("\nText: " + quotes[i].Text)

	}

	// ToDo Move models to the right file
	// Move MongoDB Operations to their own file
	// Add parsers for the other files

}
