package model

type Quote struct {
	Author   string
	QuoteText string `bson:"quote"`
	Tags     string
	Category string
}

// NewQuote NewQuote a new Quote
func (q Quote) NewQuote(author string, quoteText string, tags string, category string) Quote {
	return Quote{
		Author:   author,
		QuoteText:    quoteText,
		Tags:     tags,
		Category: category,
	}
}
