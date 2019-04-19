package parser

import (
	"encoding/json"
	"fmt"
	"github.com/St0iK/go-quote-parser/dao"
	"github.com/St0iK/go-quote-parser/model"
	"io/ioutil"
	"os"
)

// basic V1 Json parser struct
type JsonV1Parser struct {
	conf map[string]string
}

// Process
func (jv1p *JsonV1Parser) Process(conf map[string]string) (string, error) {

	file, err := os.Open(conf["FILENAME"])

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

	// loop through all the quotes found in the file
	// and insert them into the data
	for i := 0; i < len(quotes); i++ {
		dao.InsertV1(quotes[i])
	}
	return "yay",nil
}

func NewJsonV1Factory(conf map[string]string) (Parser, error) {

	return &JsonV1Parser{
		conf,
	}, nil
}