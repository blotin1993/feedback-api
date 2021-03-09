package db

import (
	"context"
	"time"

	"github.com/blotin1993/feedback-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificoRegistro permite modificar el perfil del usuario*/
func ModificoRegistro(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("feedback-db")
	col := db.Collection("users")

	/* supongo que me van a enviar de a un campo a modificar a la vez,
	   por eso me fijo si lo que viene tiene valor (largo mayor a cero)
	   Creamos un mapa de interfaces para armar el registro de actualización a la bd
	   Le pongo la información que hay que modificar*/
	registro := make(map[string]interface{})

	if len(u.Name) > 0 {
		registro["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		registro["lastname"] = u.LastName
	}
	if len(u.ProfilePicture) > 0 {
		registro["profilePicture"] = u.ProfilePicture
	}
	updtString := bson.M{
		"$set": registro,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)

	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}
