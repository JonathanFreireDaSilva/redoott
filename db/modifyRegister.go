package db

import (
	"context"
	"time"

	"github.com/JonathanFreireDaSilva/redoott/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModifyRegister permite modificar el perfil de un usuario de la basededatos*/
func ModifyRegister(u models.User, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redoott")
	collUser := db.Collection("users")

	register := make(map[string]interface{})

	if len(u.Nombre) > 0 {
		register["nombre"] = u.Nombre
	}

	register["fechaNacimiento"] = u.FechaNacimiento

	if len(u.Apellidos) > 0 {
		register["apellidos"] = u.Apellidos
	}

	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}

	if len(u.Biografia) > 0 {
		register["biografia"] = u.Biografia
	}

	if len(u.Ubicacion) > 0 {
		register["ubicacion"] = u.Ubicacion
	}

	if len(u.SitioWeb) > 0 {
		register["sitioWeb"] = u.SitioWeb
	}

	updateString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	_, err := collUser.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}

	return true, nil
}
