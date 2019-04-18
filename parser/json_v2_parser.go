package parser

import (
	"encoding/json"
	"fmt"
	"github.com/St0iK/go-quote-parser/dao"
	"github.com/St0iK/go-quote-parser/model"
	"io/ioutil"
	"os"
)

//The second implementation.
type JsonV2Parser struct {
	filename string
}

func (pds *JsonV2Parser) Process(filename string) (string, error) {
	file, err := os.Open(filename)

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

	for i := 0; i < len(quotes); i++ {
		fmt.Printf("%+v\n", quotes[i])
		dao.InsertV2(quotes[i])
		fmt.Println("\nName: " + quotes[i].Author)

	}
	return "yay",nil
}

func NewJsonV2Factory(conf map[string]string) (Parser, error) {

	return &JsonV2Parser{
		conf["FILENAME"],
	}, nil
}
