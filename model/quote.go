package model

// V1 Quote struct
type QuoteV1 struct {
	Author string `bson:"Name" json:"Name"`
	Quote  string `bson:"Text" json:"Text"`
	Tags string `bson:"Tags" json:"Tags"`
	Category string `bson:"Category" json:"Category"`
}

// V1 Quote struct
type QuoteV2 struct {
	Author string `bson:"Name" json:"quoteAuthor"`
	Quote  string `bson:"Text" json:"quoteText"`
	Category string `bson:"Category" json:"Category"`
}

// V3 Quote struct
type QuoteV3 struct {
	Author string `bson:"Name" json:"en"`
	Quote  string `bson:"Text" json:"author"`
	Category string `bson:"Category" json:"Category"`
}