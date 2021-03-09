package db

import (
	"github.com/blotin1993/feedback-api/models"
)

/*IntentoLogin realiza el chequeo de login a la BD*/
func IntentoLogin(email string, password string) (models.User, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}
	// Ahora comparo la password que en la BD está encriptada
	// creo una variable que sea un slice de bytes
	passwordBytes := []byte(password)
	// creo otra variable con la password que tengo en la BD para el usuario
	passwordBD := []byte(usu.Password)
	// Ahora llamo a una función que compara las password
	match, _ := ComparePasswords(passwordBytes, passwordBD)

	if match != true {
		return usu, false
	}
	return usu, true
}
