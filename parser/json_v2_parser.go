package parser

import (
	"encoding/json"
	"fmt"
	"github.com/St0iK/go-quote-parser/dao"
	"github.com/St0iK/go-quote-parser/model"
	"io/ioutil"
	"os"
)

// basic V2 Json parser struct
type JsonV2Parser struct {
	conf map[string]string
}

// Process
func (pds *JsonV2Parser) Process(conf map[string]string) (string, error) {
	file, err := os.Open(conf["FILENAME"])

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var quotes []model.QuoteV2

	err = json.Unmarshal(byteValue, &quotes)
	if err != nil {
		fmt.Println(err)
	}

	// loop through all the quotes found in the file
	// and insert them into the database
	for i := 0; i < len(quotes); i++ {
		dao.InsertV2(quotes[i])
	}
	return "yay",nil
}

func NewJsonV2Factory(conf map[string]string) (Parser, error) {

	return &JsonV2Parser{
		conf,
	}, nil
}
