package routers

import (
	"net/http"

	"github.com/blotin1993/feedback-api/db"
	"gopkg.in/gomail.v2"
)

/* GetMyPassword - receive the user data from DataBase and send an Email with his current password */
func GetMyPassword(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	// if everything goes well, i search the user in the db
	profile, err := db.GetProfile(ID)
	// if err is not nil, something went wrong
	if err != nil {
		http.Error(w, "ocurrio un error a lintentar buscar el registro"+err.Error(), 400)
		return
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", "feedbackapiadm@gmail.com")
	msg.SetHeader("To", profile.Email)
	msg.SetHeader("Subject", "Get your password.")
	msg.SetBody("text/html", "Hola "+profile.Name+" "+profile.LastName+"/n"+
		"Solicitaste recuperar tu contraseña: \n ")

}

func SendEmail(from string, to string, msg string) {

}
