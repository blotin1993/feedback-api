package routers

import (
	"encoding/json"
	"fmt"
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

	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Agregar validación para msg.
	// if len(fb.message) == "" {
	// 	http.Error(w, "the message must have at least one character", 400)
	// 	return
	// }

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
