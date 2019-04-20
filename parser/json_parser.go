package parser

import (
	"encoding/json"
	"fmt"
	"github.com/St0iK/go-quote-parser/dao"
	"github.com/St0iK/go-quote-parser/model"
	"io/ioutil"
	"log"
	"os"
)

// basic V1 Json parser struct
type JsonParser struct {
	conf map[string]string
}

// Process
func (jp *JsonParser) Process(conf map[string]string) (string, error) {
	log.Printf("Starting Process")
	fmt.Printf("%v\n", conf)

	file, err := os.Open(conf["FILENAME"])

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	// Declared an empty interface of type Array
	var results []map[string]interface{}

	err = json.Unmarshal(byteValue, &results)
	if err != nil {
		fmt.Println(err)
	}

	for key, result := range results {
		fmt.Println("Reading Value for Key :", key)
		fmt.Println("Inserting Quote from :", result[conf["Author"]])
		fmt.Println("Inserting Quote in category :", result[conf["Category"]])
		quote := model.Quote{
			Author: fmt.Sprint(result[conf["Author"]]),
			QuoteText: fmt.Sprint(result[conf["QuoteText"]]),
			Tags: fmt.Sprint(result[conf["Tags"]]),
			Category: fmt.Sprint(result["Category"]),
		}

		_ = dao.Insert(quote)

	}
	// update
	return "yay",nil
}

func NewJsonFactory(conf map[string]string) (Parser, error) {

	return &JsonParser{
		conf,
	}, nil
}
