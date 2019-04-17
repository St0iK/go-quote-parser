package main

import (
	"fmt"

	"github.com/St0iK/go-quote-parser/dao"
	"github.com/St0iK/go-quote-parser/parser"
)

func init() {
	dao.Connect()
}

func init() {
	parser.Register("json-v1", parser.NewJsonV1Factory)
	parser.Register("json-v2", parser.NewJsonV2Factory)
}

func main() {
	// Create a list of the files to be parsed
	// Loop through all of them
	// spawn go-routine to process the items
	// Wait all routines to finish and processing the items
	fmt.Println("Initialising parsing...")

	parser, _ := parser.CreateParser(map[string]string{
		"PARSER_NAME": "json-v1",
		"FILENAME": "quotes_v1.json",
	})
	var res, _ = parser.Process("sdfsdfsdf")
	fmt.Println(res)
}
