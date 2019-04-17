package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/St0iK/go-quote-parser/dao"
	"github.com/St0iK/go-quote-parser/model"
)
//The first implementation.
type JsonV1Parser struct {
	filename string
}

func (jv1p *JsonV1Parser) Process(filename string) (string, error) {
	
	jsonFile, err := os.Open("quotes/quotes_v1.json")

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
		fmt.Println("\nName: " + quotes[i].Name)

	}
	return "yay",nil
}

func NewJsonV1Factory(conf map[string]string) (Parser, error) {

	return &JsonV1Parser{
		"something",
	}, nil
}
