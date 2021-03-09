package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
)

//FeedbackTry .
func FeedbackTry(w http.ResponseWriter, r *http.Request) {
	rID := r.URL.Query().Get("target_id")
	if len(rID) < 1 {
		http.Error(w, "Falta el target.......", http.StatusBadRequest)
		return
	}
	var fb models.Feedback

	// decodificamos el body y armamos un registro
	err := json.NewDecoder(r.Body).Decode(&fb)
	//
	fb.IssuerID = IDUsuario
	fb.ReceiverID = rID
	fb.Date = time.Now()
	// Para insertarlo en la base de datos necesitamos mapearlo a un bson
	//-------Feedback validation------
	//TeamArea validation
	if len(fb.TeamArea.Message) == 0 {
		http.Error(w, "the message must have at least one character", 400)
		return
	}
	if len(fb.TeamArea.Message) >= 500 {
		http.Error(w, "must have a maximum of 500 characters", 400)
		return
	}

	if len(fb.TeamArea.TeamPlayer) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	if len(fb.TeamArea.Commited) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	if len(fb.TeamArea.Communication) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	//validation TechArea
	if len(fb.TechArea.Message) == 0 {
		http.Error(w, "the message must have at least one character", 400)
		return
	}

	if len(fb.TechArea.Message) >= 500 {
		http.Error(w, "must have a maximum of 500 characters", 400)
		return
	}

	if len(fb.TechArea.BestPractices) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	if len(fb.TechArea.TechKnowledge) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	if len(fb.TechArea.CodingStyle) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	//validation PerformanceArea
	if len(fb.PerformanceArea.Message) == 0 {
		http.Error(w, "the message must have at least one character", 400)
		return
	}

	if len(fb.PerformanceArea.Message) >= 500 {
		http.Error(w, "must have a maximum of 500 characters", 400)
		return
	}

	if len(fb.PerformanceArea.Message) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	if len(fb.PerformanceArea.Message) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	if len(fb.PerformanceArea.Message) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	//Validation Message
	if len(fb.Message) == 0 {
		http.Error(w, "the message must have at least one character", 400)
		return
	}

	if len(fb.Message) >= 1500 {
		http.Error(w, "must have a maximum of 1500 characters", 400)
		return
	}

	_, status, err := db.InsertoFeedback(fb)

	//si hay un error
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, intentelo nuevamente.  "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet.", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
