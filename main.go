package main

import (
	"fmt"
	"github.com/St0iK/go-quote-parser/dao"
	"github.com/St0iK/go-quote-parser/parser"
	"io/ioutil"
	"log"
)

func init() {
	dao.Connect()
}

func init() {
	parser.Register("quotes_v1.json", parser.NewJsonV1Factory)
	parser.Register("quotes_v2.json", parser.NewJsonV2Factory)
}

func main() {
	// Create a list of the files to be parsed
	quotesFolder := "./quotes"
	files, err := ioutil.ReadDir(quotesFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		file := quotesFolder + "/" + f.Name()
		parser, _ := parser.CreateParser(map[string]string{
			"FILENAME": file,
		})
		if parser != nil {
			var res, _ = parser.Process(file)
			fmt.Println(res)
		}
	}






	// Loop through all of them
	// spawn go-routine to process the items
	// Wait all routines to finish and processing the items
	//fmt.Println("Initialising parsing...")
	//
	//parser, _ := parser.CreateParser(map[string]string{
	//	"PARSER_NAME": "json-v1",
	//	"FILENAME": "quotes_v1.json",
	//})
	//var res, _ = parser.Process("sdfsdfsdf")
	//fmt.Println(res)
}
