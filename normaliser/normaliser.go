package normaliser

import (
	"github.com/St0iK/go-quote-parser/model"
	"log"
)

type Normaliser struct {
}

// Normalise the article
func (n *Normaliser) Normalise(author string, quote string, category string) (model.Quote, error) {

	q := model.Quote{}.NewQuote(author, quote, "", category)

	log.Printf("Normalised data %s", q)
	return q, nil
}
