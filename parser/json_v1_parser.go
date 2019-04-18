package parser

import (
	"encoding/json"
	"fmt"
	"github.com/St0iK/go-quote-parser/dao"
	"github.com/St0iK/go-quote-parser/model"
	"io/ioutil"
	"os"
)

//The first implementation.
type JsonV1Parser struct {
	filename string
}

func (jv1p *JsonV1Parser) Process(filename string) (string, error) {
	
	file, err := os.Open(filename)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var quotes []model.QuoteV1

	err = json.Unmarshal(byteValue, &quotes)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(quotes); i++ {
		fmt.Printf("%+v\n", quotes[i])
		dao.InsertV1(quotes[i])
		fmt.Println("\nName: " + quotes[i].Author)
		fmt.Println("TTTT: " + quotes[i].Quote)

	}
	return "yay",nil
}

func NewJsonV1Factory(conf map[string]string) (Parser, error) {

	return &JsonV1Parser{
		conf["FILENAME"],
	}, nil
}
