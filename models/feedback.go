package models
 
import (
    "go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Feedback struct {
    
    IssuerID          			primitive.ObjectID  `bson:"issuer_id,omitempty" json:"issuer_id"`
	ReceiverID          		primitive.ObjectID  `bson:"receiver_id,omitempty" json:"receiver_id"`
    Date 						time.Time 			`bson:"date" json:"date,omitempty"`	
	TechArea					struct {
		Message string              `bson:"message,omitempty" json:"message"`
		TechKnowledge string 		`bson:"techKnowledge,omitempty" json:"techKnowledge"`
		BestPractices string 		`bson:"bestPractices,omitempty" json:"bestPractices"`
		CodingStyle string 			`bson:"codingStyle,omitempty" json:"codingStyle"`
	}
	TeamArea					struct {
		Message string              `bson:"message,omitempty" json:"message"`
		TeamPlayer string 			`bson:"teamPlayer,omitempty" json:"teamPlayer"`
		Commited string 			`bson:"commited,omitempty" json:"commited"`
		Communication string 		`bson:"communication,omitempty" json:"communication"`
	}
	PerformanceArea					struct {
		Message string              `bson:"message,omitempty" json:"message"`
		WorkQuality string 			`bson:"workQuality,omitempty" json:"workQuality"`
		ClientOriented string 		`bson:"clientOriented,omitempty" json:"clientOriented"`
	}
}
