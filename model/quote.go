package model

type Quote struct {
	Author   string
	QuoteText    string
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

// V1 Quote struct
type QuoteV1 struct {
	Author   string `bson:"Name" json:"Name"`
	Quote    string `bson:"Text" json:"Text"`
	Tags     string `bson:"Tags" json:"Tags"`
	Category string `bson:"Category" json:"Category"`
}

// V1 Quote struct
type QuoteV2 struct {
	Author   string `bson:"Name" json:"quoteAuthor"`
	Quote    string `bson:"Text" json:"quoteText"`
	Category string `bson:"Category" json:"Category"`
}

// V3 Quote struct
type QuoteV3 struct {
	Author   string `bson:"Name" json:"en"`
	Quote    string `bson:"Text" json:"author"`
	Category string `bson:"Category" json:"Category"`
}
