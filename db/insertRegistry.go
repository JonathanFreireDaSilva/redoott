package db

import (
	"context"
	"time"

	"github.com/JonathanFreireDaSilva/redoott/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertRegistry es la parada final con la DB para insertar los datos del usuario*/
func InsertRegistry(user models.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //le decimos al context.TODO que teniamos que espere 15 segundos
	defer cancel()

	db := MongoCN.Database("redoott")
	userCollection := db.Collection("users")

	user.Password, _ = EncryptPassword(user.Password)

	res, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := res.InsertedID.(primitive.ObjectID) // manera de obtengo el id
	return ObjID.String(), true, nil
}
