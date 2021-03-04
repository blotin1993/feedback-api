package models
 
import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Metrics struct {
	Id 			primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Text 		string     			`bson:"text,omitempty" json:"text"`
}
























/*
Métricas de Feedback:
 1 - Let´s work on this.
 2 - Reach the Goal.
 3 - Relevant Performance.
 4 - Master.
*/