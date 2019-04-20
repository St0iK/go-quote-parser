package main

import (
	"fmt"
	"github.com/St0iK/go-quote-parser/dao"
	"github.com/St0iK/go-quote-parser/parser"
	"io/ioutil"
	"log"
)

func init() {
	// Connect to the database
	dao.Connect()

	// register the parser factories
	parser.Register("quotes_v1.json", parser.NewJsonFactory)
	parser.Register("quotes_v2.json", parser.NewJsonFactory)
	parser.Register("quotes_v3.json", parser.NewJsonFactory)
}

func main() {

	// TODO:
	//  1. move configuration into a file
	//  2. add logger and handle errors properly

	// folder containing the quotes files
	qf := "./quotes"

	var config = map[string]map[string]string{}
	config["quotes_v1.json"] = map[string]string{
		"FILENAME":  qf + "/quotes_v1.json",
		"Author":    "Name",
		"QuoteText": "Text",
	}

	config["quotes_v2.json"] = map[string]string{
		"FILENAME":  qf + "/quotes_v2.json",
		"Author":    "quoteAuthor",
		"QuoteText": "quoteText",
	}

	config["quotes_v3.json"] = map[string]string{
		"FILENAME":  qf + "/quotes_v3.json",
		"Author":    "author",
		"QuoteText": "en",
	}

	files, err := ioutil.ReadDir(qf)
	if err != nil {
		log.Fatal(err)
	}

	// loop through the files in the folder
	for _, f := range files {
		// crate a parser and pass the filename
		p, _ := parser.GetParserForFile(config[f.Name()])

		// If parser was found for this file, then process it
		if p != nil {
			var res, _ = p.Process()
			fmt.Println(res)
		}
	}
}
