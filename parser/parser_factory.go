package parser

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
)

// ParserFactory ...
type ParserFactory func(conf map[string]string) (Parser, error)

var parserFactories = make(map[string]ParserFactory)

// Register ...
func Register(name string, factory ParserFactory) {
	if factory == nil {
		log.Panicf("Parser factory %s does not exist.", name)
	}
	_, registered := parserFactories[name]
	if registered {
		log.Panicf("Datastore factory %s already registered. Ignoring.", name)
	}
	parserFactories[name] = factory
}

func CreateParser(conf map[string]string) (Parser, error) {

	parserName := filepath.Base(conf["FILENAME"])

	parserFactory, ok := parserFactories[parserName]
	if !ok {

		return nil, errors.New(fmt.Sprintf("Invalid Datastore name. Must be one of"))
	}

	// Run the factory with the configuration.
	return parserFactory(conf)
}
