package models
 
import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Metrics struct {
	Id 			primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Text 		string     			`bson:"text,omitempty" json:"text"`
}
