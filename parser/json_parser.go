package parser

import (
	"encoding/json"
	"fmt"
	"github.com/St0iK/go-quote-parser/config"
	"github.com/St0iK/go-quote-parser/dao"
	"github.com/St0iK/go-quote-parser/model"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// basic V1 Json parser struct
type JsonParser struct {
	conf config.ParserConfig
}

func empty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// Process
func (jp *JsonParser) Process() (string, error) {
	log.Printf("Starting json parser")
	fmt.Printf("Configuration: %v\n", jp.conf)

	file, err := os.Open(jp.conf.File)

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
		log.Println("Reading Value for Key :", key)
		log.Println("Inserting Quote from :", result[jp.conf.Author])
		log.Println("Inserting Quote: :", result[jp.conf.QuoteText])

		var author = fmt.Sprint(result[jp.conf.Author])
		var quoteText = fmt.Sprint(result[jp.conf.QuoteText])
		var tags = fmt.Sprint(result[jp.conf.Tags])
		var category = fmt.Sprint(result[jp.conf.Category])

		if !empty(author) && !empty(quoteText) {
			quote := model.Quote{
				Author:    author,
				QuoteText: quoteText,
				Tags:      tags,
				Category:  category,
			}
			dao.Insert(quote)
		}

	}

	return "yay", nil
}

func NewJsonFactory(conf config.ParserConfig) (Parser, error) {

	return &JsonParser{
		conf,
	}, nil
}
