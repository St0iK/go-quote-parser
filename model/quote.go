package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Quote struct
type Quote struct {
	ID   primitive.ObjectID `bson:"_id" json:"-"`
	Name string             `bson:"Name" json:"Name"`
	Text string             `bson:"Text" json:"Text"`
}
