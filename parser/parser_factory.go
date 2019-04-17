package parser

import (
	"errors"
	"fmt"
	"log"
	"strings"
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

	// Query configuration for datastore defaulting to "memory".
	engineName := conf["PARSER_NAME"]

	engineFactory, ok := parserFactories[engineName]
	if !ok {
		// Factory has not been registered.
		// Make a list of all available datastore factories for logging.
		availableDatastores := make([]string, len(parserFactories))
		for k, _ := range parserFactories {
			availableDatastores = append(availableDatastores, k)
		}
		return nil, errors.New(fmt.Sprintf("Invalid Datastore name. Must be one of: %s", strings.Join(availableDatastores, ", ")))
	}

	// Run the factory with the configuration.
	return engineFactory(conf)
}
