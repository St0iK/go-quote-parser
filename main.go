package main

import (
    "fmt"
    "github.com/St0iK/go-quote-parser/config"
    "github.com/St0iK/go-quote-parser/dao"
    "github.com/St0iK/go-quote-parser/parser"
    "log"
    "os"
)

var c config.Configuration

func init() {
    // Connect to the database
    dao.Connect()

    // Register the parser factories
    parser.Register("V1", parser.NewJsonFactory)
    parser.Register("V2", parser.NewJsonFactory)
    parser.Register("V3", parser.NewJsonFactory)

    // Load configuration
    if err := c.GetConf(); err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }
}

func main() {
    log.Println("Initializing Go-Quote-Parser")

    log.Println("Reading all files inside quotes folder")
    files, err := os.ReadDir(c.QuotesFolder)
    if err != nil {
        log.Fatalf("Failed to read directory %s: %v", c.QuotesFolder, err)
    }

    // Loop through the files in the folder
    for _, f := range files {
        parserConfig := c.GetConfForFile(f.Name())

        // Create a parser and pass the filename
        p, err := parser.GetParserForFile(parserConfig)
        if err != nil {
            log.Printf("Failed to get parser for file %s: %v", f.Name(), err)
            continue
        }

        // If parser was found for this file, process it
        if p != nil {
            res, err := p.Process()
            if err != nil {
                log.Printf("Error processing file %s: %v", f.Name(), err)
                continue
            }
            fmt.Println(res)
        }
    }
}
