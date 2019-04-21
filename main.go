package main

import (
	"fmt"
	"github.com/St0iK/go-quote-parser/config"
	"github.com/St0iK/go-quote-parser/dao"
	"github.com/St0iK/go-quote-parser/parser"
	"io/ioutil"
	"log"
)

var c config.Configuration

func init() {
	// Connect to the database
	dao.Connect()

	// register the parser factories
	parser.Register("V1", parser.NewJsonFactory)
	parser.Register("V2", parser.NewJsonFactory)
	parser.Register("V3", parser.NewJsonFactory)

	c.GetConf()
}

func main() {
	log.Println("Initialising Go-Quote-Parser")

	log.Println("Reading all files inside quotes folder")
	files, err := ioutil.ReadDir(c.QuotesFolder)
	if err != nil {
		log.Fatal(err)
	}

	// loop through the files in the folder
	for _, f := range files {
		parserConfig := c.GetConfForFile(f.Name())
		// crate a parser and pass the filename
		p, _ := parser.GetParserForFile(parserConfig)

		// If parser was found for this file, then process it
		if p != nil {
			var res, _ = p.Process()
			fmt.Println(res)
		}
	}
}
