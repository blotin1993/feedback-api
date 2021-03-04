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





/*
Estructura de un Feedback:
Emisor: Persona que formula el feedback
Receptor: Persona que recibe el feedback
Fecha: en la cual fue completado un feedback

Area Técnica:  Texto 500.
 Conocimiento de una Tecnologia. MF
 Buenas Prácticas. MF
 Estilo de Codeo. MF

Area de Equipo: texto 500.
 Team Player: Trabajo en equipo MF
 Commited. MF
 Communication. MF 

Area de Desempeño: Texto 500 
 Calidad del Trabajo MF
 Orientado/a al Cliente: MF

Comments: texto,1500 
 
Métricas de Feedback:
 1 - Let´s work on this.
 2 - Reach the Goal.
 3 - Relevant Performance.
 4 - Master.

Nota: completar al menos un área de desempeño.


*/