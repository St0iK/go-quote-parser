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

	// folder containing the quotes files
	quotesFolder := "./quotes"

	var configuration = map[string]map[string]string{}
	configuration["quotes_v1.json"] = map[string]string{
		"FILENAME": quotesFolder + "/quotes_v1.json",
		"Author":"Name",
		"QuoteText":"Text",
	}

	configuration["quotes_v2.json"] = map[string]string{
		"FILENAME": quotesFolder + "/quotes_v2.json",
		"Author":"quoteAuthor",
		"QuoteText":"quoteText",
	}

	configuration["quotes_v3.json"] = map[string]string{
		"FILENAME": quotesFolder + "/quotes_v3.json",
		"Author":"author",
		"QuoteText":"en",
	}



	files, err := ioutil.ReadDir(quotesFolder)
	if err != nil {
		log.Fatal(err)
	}

	// loop through the files in the folder
	for _, f := range files {
		// build filename
		file := quotesFolder + "/" + f.Name()

		// crate a parser and pass the filename
		parser, _ := parser.GetParserForFile(map[string]string{
			"FILENAME": file,
		})
		// If parser was found for this file, then process it
		if parser != nil {
			var res, _ = parser.Process(configuration[f.Name()])
			fmt.Println(res)
		}
	}
}
